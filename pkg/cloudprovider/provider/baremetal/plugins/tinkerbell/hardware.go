package tinkerbell

import (
	"encoding/json"

	"github.com/tinkerbell/tink/pkg"

	"k8s.io/klog"
)

type Hardware struct {
	Spec pkg.HardwareWrapper
}

func (h *Hardware) GetName() string {
	return ""
}

func (h *Hardware) GetID() string {
	return h.Spec.Id
}

func (h *Hardware) GetIPAddress() string {
	if len(h.Spec.Network.Interfaces) > 0 {
		return h.Spec.Network.Interfaces[0].Dhcp.Ip.Address
	}

	return ""
}

func (h *Hardware) GetMACAddress() string {
	if len(h.Spec.Network.Interfaces) > 0 {
		return h.Spec.Network.Interfaces[0].Dhcp.Mac
	}

	return ""
}

func (h *Hardware) GetStatus() string {
	metadata := struct {
		State string `json:"state"`
	}{}

	if err := json.Unmarshal([]byte(h.Spec.Metadata), &metadata); err != nil {
		klog.Errorf("failed to unmarshal hardware metadata: %v", err)
		return ""
	}

	return metadata.State
}
