package config

type Config struct {
	Debug    bool
	LogLevel string

	Portals []*PortalDef
}

type PortalDef struct {
	Name    *string
	Command *string
	Hint    *string

	Aws *PortalAwsConfig
}

type PortalAwsConfig struct {
	InstanceId *string `mapstructure:"instance-id"`
	Region     *string
}
