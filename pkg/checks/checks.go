package checks

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"

	acsapi "github.com/openshift/openshift-azure/pkg/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	appclient "k8s.io/client-go/kubernetes/typed/apps/v1"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-05-01/network"
)

// WaitForHTTPStatusOk poll until URL returns 200
func WaitForHTTPStatusOk(ctx context.Context, transport http.RoundTripper, urltocheck string) error {
	cli := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	req, err := http.NewRequest("GET", urltocheck, nil)
	if err != nil {
		return err
	}
	return wait.PollUntil(time.Second, func() (bool, error) {
		resp, err := cli.Do(req)
		if err, ok := err.(*url.Error); ok {
			if err, ok := err.Err.(*net.OpError); ok {
				if err, ok := err.Err.(*os.SyscallError); ok {
					if err.Err == syscall.ENETUNREACH {
						return false, nil
					}
				}
			}
			if err.Timeout() || err.Err == io.EOF || err.Err == io.ErrUnexpectedEOF {
				return false, nil
			}
		}
		if err == io.EOF {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		return resp != nil && resp.StatusCode == http.StatusOK, nil
	}, ctx.Done())
}

func checkNamespace(namespace string) bool {
	if namespace == "default" {
		return true
	}

	whitelistedNamespaces := []string{
		"openshift-",
		"kube-",
	}
	for _, ns := range whitelistedNamespaces {
		if strings.HasPrefix(namespace, ns) {
			return true
		}
	}
	return false
}

// WaitForInfraServices verify daemonsets, statefulsets
func WaitForInfraServices(ctx context.Context, appclient *appclient.AppsV1Client) error {

	// Daemonsets
	dsList, err := appclient.DaemonSets("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, ds := range dsList.Items {
		if !checkNamespace(ds.Namespace) {
			continue
		}
		for {
			ds, err := appclient.DaemonSets(ds.Namespace).Get(ds.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
			if ds.Status.NumberMisscheduled > 0 {
				return fmt.Errorf("Daemonset[%v] in Namespace[%v] has missscheduled[%v]", ds.Name, ds.Namespace, ds.Status.NumberMisscheduled)
			}

			if ds.Status.DesiredNumberScheduled == ds.Status.CurrentNumberScheduled &&
				ds.Status.DesiredNumberScheduled == ds.Status.NumberReady &&
				ds.Status.DesiredNumberScheduled == ds.Status.UpdatedNumberScheduled &&
				ds.Generation == ds.Status.ObservedGeneration {
				break
			}
			select {
			case <-time.After(2 * time.Second):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	// Statefulsets
	ssList, err := appclient.StatefulSets("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, ss := range ssList.Items {
		if !checkNamespace(ss.Namespace) {
			continue
		}
		for {
			ss, err := appclient.StatefulSets(ss.Namespace).Get(ss.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
			if ss.Status.Replicas == ss.Status.ReadyReplicas &&
				ss.Status.ReadyReplicas == ss.Status.CurrentReplicas &&
				ss.Spec.Replicas != nil &&
				*ss.Spec.Replicas == ss.Status.Replicas &&
				ss.Generation == ss.Status.ObservedGeneration {
				break
			}
			select {
			case <-time.After(2 * time.Second):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	// Deployments
	dList, err := appclient.Deployments("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, d := range dList.Items {
		if !checkNamespace(d.Namespace) {
			continue
		}
		for {
			d, err := appclient.Deployments(d.Namespace).Get(d.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
			if d.Status.Replicas == d.Status.ReadyReplicas &&
				d.Status.ReadyReplicas == d.Status.AvailableReplicas &&
				d.Spec.Replicas != nil &&
				*d.Spec.Replicas == d.Status.Replicas &&
				d.Status.Replicas == d.Status.UpdatedReplicas &&
				d.Generation == d.Status.ObservedGeneration {
				break
			}
			select {
			case <-time.After(2 * time.Second):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	return nil
}

func CheckDNS(ctx context.Context, eipc network.PublicIPAddressesClient, lbc network.LoadBalancerFrontendIPConfigurationsClient, cs *acsapi.ContainerService) error {

	// get master FQDN from ip-apiserver loadbalancer
	ip, err := eipc.Get(ctx, cs.Properties.AzProfile.ResourceGroup, "ip-apiserver", "")
	if err != nil {
		return err
	}

	if *ip.DNSSettings.Fqdn != cs.Properties.FQDN {
		return fmt.Errorf("ip-address fqdn %s do not match configuration %s", *ip.DNSSettings.Fqdn, cs.Properties.FQDN)
	}

	// check router FQDN by quering k8s LoadBalancer and ipAddress associated with it
	lbIPList, err := lbc.List(ctx, cs.Properties.AzProfile.ResourceGroup, "kubernetes")
	if err != nil {
		return err
	}
	// iterate over k8s LoadBalancer and check if we have a match. Match should always exist, so failed by default.
	// TODO: With multiple routers this need to check multiple conditions
	exist := false
	for _, v := range lbIPList.Values() {
		parts := strings.Split(*v.PublicIPAddress.ID, "/")
		ipAddressName := parts[len(parts)-1]
		ip, err := eipc.Get(ctx, cs.Properties.AzProfile.ResourceGroup, ipAddressName, "")
		if err != nil {
			return err
		}

		if *ip.DNSSettings.Fqdn == cs.Properties.OrchestratorProfile.OpenShiftConfig.RouterProfiles[0].FQDN {
			exist = true
		}
	}
	if !exist {
		return fmt.Errorf("none of the ipAddress fqdn associated with kubernetes loadbalancer matches router fqdn from the config")
	}

	return nil
}
