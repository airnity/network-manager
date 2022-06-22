package config

type Config struct {
	Tunnels []Tunnel   `mapstructure:"tunnels" validate:"required,dive"`
	Logs    *LogConfig `mapstructure:"logs"`
}

type Tunnel struct {
	Name   string `mapstructure:"name" validate:"required"`
	Remote string `mapstructure:"remote" validate:"required,ipv4"`
	Local  string `mapstructure:"local" validate:"required,ipv4"`
	Addr   string `mapstructure:"addr" validate:"required,cidrv4"`
	State  string `mapstructure:"state"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}
