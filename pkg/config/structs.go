package config

type Config struct {
	Debug    bool
	LogLevel string

	DefaultPublicKey  string `mapstructure:"default-public-key"`
	DefaultPrivateKey string `mapstructure:"default-private-key"`

	Portals []*PortalDef
}

type PortalDef struct {
	Name *string
	Hint *string

	Aws *PortalAwsConfig
	Raw *PortalRawConfig
}

type PortalRawConfig struct {
	Command *string
}

type PortalAwsConfig struct {
	InstanceId *string `mapstructure:"instance-id"`
	Region     *string
	User       *string
}
