package plugins

// PluginDriver manages the communications between the machine controller cloud provider and the bare metal env.
type PluginDriver interface {
	GetServer(serverId string) (Server, error)
	ListServers() ([]Server, error)
	ProvisionServer() (string, error)
	DeprovisionServer() error
}

// Server represents the server/instance which exists in the bare metal env.
type Server interface {
	GetName() string
	GetID() string
	GetIPAddress() string
	GetMACAddress() string
	GetStatus() string
}
