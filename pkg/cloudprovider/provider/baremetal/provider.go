package baremetal

import (
	"github.com/kubermatic/machine-controller/pkg/apis/cluster/v1alpha1"
	"github.com/kubermatic/machine-controller/pkg/cloudprovider/instance"
	"github.com/kubermatic/machine-controller/pkg/cloudprovider/provider/baremetal/plugins"
	cloudprovidertypes "github.com/kubermatic/machine-controller/pkg/cloudprovider/types"
	"github.com/kubermatic/machine-controller/pkg/providerconfig"
	"k8s.io/klog"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

type bareMetalServer struct {
	server plugins.Server
}

func (b bareMetalServer) Name() string {
	return b.server.GetName()
}

func (b bareMetalServer) ID() string {
	return b.server.GetID()
}

func (b bareMetalServer) Addresses() map[string]corev1.NodeAddressType {
	return map[string]corev1.NodeAddressType{
		b.server.GetIPAddress(): corev1.NodeInternalIP,
	}
}

func (b bareMetalServer) Status() instance.Status {
	return instance.StatusRunning
}

type provider struct {
	configVarResolver *providerconfig.ConfigVarResolver

	plugin plugins.PluginDriver
}

// New returns a BareMetal provider
func New(configVarResolver *providerconfig.ConfigVarResolver, driver plugins.PluginDriver) cloudprovidertypes.Provider {
	if driver == nil {
		klog.Error("No bare metal plugin driver was provided")
		return nil
	}

	return &provider{
		configVarResolver: configVarResolver,
		plugin:            driver,
	}
}

func (p provider) AddDefaults(spec v1alpha1.MachineSpec) (v1alpha1.MachineSpec, error) {
	return spec, nil
}

func (p provider) Validate(machinespec v1alpha1.MachineSpec) error {
	return nil
}

func (p provider) Get(machine *v1alpha1.Machine, data *cloudprovidertypes.ProviderData) (instance.Instance, error) {
	panic("implement me")
}

func (p provider) GetCloudConfig(spec v1alpha1.MachineSpec) (config string, name string, err error) {
	return "", "", nil
}

func (p provider) Create(machine *v1alpha1.Machine, data *cloudprovidertypes.ProviderData, userdata string) (instance.Instance, error) {
	panic("implement me")
}

func (p provider) Cleanup(machine *v1alpha1.Machine, data *cloudprovidertypes.ProviderData) (bool, error) {
	panic("implement me")
}

func (p provider) MachineMetricsLabels(machine *v1alpha1.Machine) (map[string]string, error) {
	return nil, nil
}

func (p provider) MigrateUID(machine *v1alpha1.Machine, new types.UID) error {
	return nil
}

func (p provider) SetMetricsForMachines(machines v1alpha1.MachineList) error {
	return nil
}
