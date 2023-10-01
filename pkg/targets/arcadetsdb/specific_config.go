package arcadetsdb

import "github.com/blagojts/viper"

type SpecificConfig struct {
	host string
	port int
}

func parseSpecificConfig(v *viper.Viper) *SpecificConfig {
	return &SpecificConfig{
		host: v.GetString("host"),
		port: v.GetInt("port"),
	}
}
