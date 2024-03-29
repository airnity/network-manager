package config

type Config struct {
	VRFs     []VRF      `mapstructure:"vrfs" validate:"required,dive"`
	Tunnels  []Tunnel   `mapstructure:"tunnels" validate:"required,dive"`
	NatRules NatRules   `mapstructure:"natRules" validate:"required"`
	Logs     *LogConfig `mapstructure:"logs"`
}

type VRF struct {
	Name    string `mapstructure:"name" validate:"required"`
	TableID int    `mapstructure:"tableID" validate:"required"`
}

type Tunnel struct {
	Name   string `mapstructure:"name" validate:"required"`
	Remote string `mapstructure:"remote" validate:"required,ipv4"`
	Local  string `mapstructure:"local" validate:"required,ipv4"`
	Addr   string `mapstructure:"addr" validate:"required,cidrv4"`
	VRF    string `mapstructure:"vrf"`
	State  string `mapstructure:"state"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type NatRules struct {
	DestNat   []NatRule `mapstructure:"destNat" validate:"required,dive"`
	SourceNat []NatRule `mapstructure:"sourceNat" validate:"required,dive"`
}

type NatRule struct {
	Source       string `mapstructure:"source"`
	Destination  string `mapstructure:"destination" validate:"required"`
	Interface    string `mapstructure:"interface" validate:"required"`
	TranslatedIP string `mapstructure:"translatedIP" validate:"required"`
	Port         int    `mapstructure:"port"`
	Proto        string `mapstructure:"proto"`
}
