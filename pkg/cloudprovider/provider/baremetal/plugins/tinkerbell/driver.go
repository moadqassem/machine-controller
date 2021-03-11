package tinkerbell

import (
	"context"
	"errors"
	"fmt"
	"github.com/kubermatic/machine-controller/pkg/cloudprovider/provider/baremetal/plugins"
	tinkerbellclient "github.com/kubermatic/machine-controller/pkg/cloudprovider/provider/baremetal/plugins/tinkerbell/client"
	tinkclient "github.com/tinkerbell/tink/client"
	"github.com/tinkerbell/tink/pkg"
	"github.com/tinkerbell/tink/protos/template"
	"gopkg.in/yaml.v3"
)

type driver struct {
	tinkServerAddress string
	imageRepoAddress  string

	hardwareClient *tinkerbellclient.Hardware
	templateClient *tinkerbellclient.Template
	workflowClient *tinkerbellclient.Workflow
}

// NewTinkerbellDriver returns a new TinkerBell driver with a configured tinkserver address and a client timeout.
func NewTinkerbellDriver(tinkServerAddress, imageRepoAddress string) (*driver, error) {
	if tinkServerAddress == "" || imageRepoAddress == "" {
		return nil, errors.New("tink-server address, imageRepoAddress cannot be empty")
	}

	if err := tinkclient.Setup(); err != nil {
		return nil, fmt.Errorf("failed to setup tink-server client: %v", err)
	}

	d := &driver{
		tinkServerAddress: tinkServerAddress,
		imageRepoAddress:  imageRepoAddress,
		hardwareClient:    tinkerbellclient.NewHardwareClient(tinkclient.HardwareClient),
		workflowClient:    tinkerbellclient.NewWorkflowClient(tinkclient.WorkflowClient, tinkerbellclient.NewHardwareClient(tinkclient.HardwareClient)),
		templateClient:    tinkerbellclient.NewTemplateClient(tinkclient.TemplateClient),
	}

	return d, nil
}
func (d *driver) GetServer(ctx context.Context, serverId, macAddress, ipAddress string) (plugins.Server, error) {
	if serverId == "" {
		return nil, errors.New("server id cannot be empty")
	}

	hw, err := d.hardwareClient.Get(ctx, serverId, ipAddress, macAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get hardware: %v", err)
	}

	return &Hardware{
		Spec: pkg.HardwareWrapper{
			hw,
		},
	}, nil
}

func (d *driver) ProvisionServer(ctx context.Context, server plugins.Server) (string, error) {
	hw := server.(*Hardware).Spec.Hardware
	if err := d.hardwareClient.Create(ctx, hw); err != nil {
		return "", fmt.Errorf("failed to register hardware to tink-server: %v", err)
	}

	tmpl := createTemplate(server.GetMACAddress(), d.tinkServerAddress, d.imageRepoAddress)
	payload, err := yaml.Marshal(tmpl)
	if err != nil {
		return "", fmt.Errorf("failed marshalling workflow template: %v", err)
	}

	workflowTemplate := &template.WorkflowTemplate{
		Name: tmpl.Name,
		Data: string(payload),
	}

	if err := d.templateClient.Create(ctx, workflowTemplate); err != nil {
		return "", fmt.Errorf("failed to create workflow template: %v", err)
	}

	return d.workflowClient.Create(ctx, workflowTemplate.Id, server.GetID())
}

func (d *driver) DeprovisionServer(serverId string) (string, error) {
	return "", nil
}
