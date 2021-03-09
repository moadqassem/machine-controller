package tinkerbell

type Hardware struct {
	ID       string   `json:"id"`
	Metadata Metadata `json:"metadata"`
	Network  Network  `json:"network"`
}

func (t *Hardware) GetName() string {
	return ""
}

func (t *Hardware) GetID() string {
	return t.ID
}

func (t *Hardware) GetIPAddress() string {
	if len(t.Network.Interfaces) > 0 {
		return t.Network.Interfaces[0].DHCP.IP.Address
	}

	return ""
}

func (t *Hardware) GetMACAddress() string {
	if len(t.Network.Interfaces) > 0 {
		return t.Network.Interfaces[0].DHCP.Mac
	}

	return ""
}

func (t *Hardware) GetStatus() string {
	return t.Metadata.State
}

type Metadata struct {
	Facility Facility `json:"facility"`
	State    string   `json:"state"`
}

type Facility struct {
	FacilityCode string `json:"facility_code"`
	PlanSlug     string `json:"plan_slug"`
}

type Network struct {
	Interfaces []Interface `json:"interfaces"`
}

type Interface struct {
	DHCP    DHCP    `json:"dhcp"`
	NetBoot NetBoot `json:"net_boot"`
}

type DHCP struct {
	Arch string `json:"arch"`
	IP   IP     `json:"ip"`
	Mac  string `json:"mac"`
	UEFI bool   `json:"uefi"`
}

type IP struct {
	Address string `json:"address"`
	Gateway string `json:"gateway"`
	NetMask string `json:"netmask"`
}

type NetBoot struct {
	AllowPXE      bool `json:"allow_pxe"`
	AllowWorkflow bool `json:"allow_workflow"`
}
