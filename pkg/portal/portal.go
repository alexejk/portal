package portal

type Portal struct {

	Name string
	Type string

	Tunnel *TunnelConfig
	Remote *RemoteConfig
}

type TunnelConfig struct {
	Local *int
	Remote *RemoteConfig
}

type RemoteConfig struct {
	Host string
	Port *int

	Credentials *RemoteCredentials
}

type RemoteCredentials struct {
	Key *string
	Username *string
}
