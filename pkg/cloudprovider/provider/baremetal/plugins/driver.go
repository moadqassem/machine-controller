package plugins

import "context"

// PluginDriver manages the communications between the machine controller cloud provider and the bare metal env.
type PluginDriver interface {
	GetServer(ctx context.Context, serverId, macAddress, ipAddress string) (Server, error)
	ProvisionServer(context.Context, Server) (string, error)
	DeprovisionServer(serverId string) (string, error)
}

// Server represents the server/instance which exists in the bare metal env.
type Server interface {
	GetName() string
	GetID() string
	GetIPAddress() string
	GetMACAddress() string
	GetStatus() string
}
