package validate

import (
	"errors"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghodss/yaml"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/api/v1"
)

/*
name: openshift
location: eastus
properties:
  openShiftVersion: "$DEPLOY_VERSION"
  publicHostname: openshift.$RESOURCEGROUP.$DNS_DOMAIN
  routerProfiles:
  - name: default
    publicSubdomain: $RESOURCEGROUP.$DNS_DOMAIN
  agentPoolProfiles:
  - name: master
    role: master
    count: 3
    vmSize: Standard_D2s_v3
    osType: Linux
  - name: infra
    role: infra
    count: 1
    vmSize: Standard_D2s_v3
    osType: Linux
  - name: compute
    role: compute
    count: 1
    vmSize: Standard_D2s_v3
    osType: Linux
  servicePrincipalProfile:
    clientID: $AZURE_CLIENT_ID
    secret: $AZURE_CLIENT_SECRET
*/

var testOpenShiftClusterYAML = []byte(`---
location: eastus
name: openshift
properties:
  openShiftVersion: v3.10
  publicHostname: openshift.test.example.com
  routerProfiles:
  - name: default
    publicSubdomain: test.example.com
  masterPoolProfile:
    name: master
    count: 3
    vmSize: Standard_D2s_v3
    osType: Linux
  agentPoolProfiles: 
  - name: infra
    role: infra
    count: 1
    vmSize: Standard_D2s_v3
    osType: Linux
  - name: compute
    role: compute
    count: 1
    vmSize: Standard_D2s_v3
    osType: Linux
  servicePrincipalProfile:
    clientID: client_id
    secret: client_secret
`)

func TestValidate(t *testing.T) {

	tests := map[string]struct {
		f            func(*v1.OpenShiftCluster)
		expectedErrs []error
	}{
		"test yaml parsing": { // test yaml parsing

		},
		"test version": {
			f:            func(oc *v1.OpenShiftCluster) { oc.Properties.OpenShiftVersion = "v3.11" },
			expectedErrs: []error{errors.New("invalid properties.openShiftVersion \"v3.11\"")},
		},
		"test Location": {
			f:            func(oc *v1.OpenShiftCluster) { oc.Location = "" },
			expectedErrs: []error{errors.New("invalid location \"\"")},
		},
		"test Name": {
			f:            func(oc *v1.OpenShiftCluster) { oc.Name = "" },
			expectedErrs: []error{errors.New("invalid name \"\"")},
		},
		"test ProvisioningState": {
			f:            func(oc *v1.OpenShiftCluster) { oc.Properties.ProvisioningState = "testing" },
			expectedErrs: []error{errors.New("invalid properties.provisioningState \"testing\"")},
		},
		"test master count": {
			f:            func(oc *v1.OpenShiftCluster) { oc.Properties.MasterPoolProfile.Count = 1 },
			expectedErrs: []error{errors.New("invalid masterPoolProfile.count 1")},
		},
	}

	for name, test := range tests {
		var oc *v1.OpenShiftCluster
		err := yaml.Unmarshal(testOpenShiftClusterYAML, &oc)
		if err != nil {
			t.Fatal(err)
		}

		if test.f != nil {
			test.f(oc)
		}

		// TODO quick fix but means we're hoping conversion is correct.
		cs := api.ConvertVLabsOpenShiftClusterToContainerService(oc)
		errs := ContainerService(cs, nil)
		if !reflect.DeepEqual(errs, test.expectedErrs) {
			t.Errorf("%s expected errors %#v but received %#v", name, spew.Sprint(test.expectedErrs), spew.Sprint(errs))
		}

	}
}
