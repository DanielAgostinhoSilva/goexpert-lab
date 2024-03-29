package env

import (
	"github.com/spf13/viper"
	"log"
)

type EnvConfig struct {
	WebServerPort            string `mapstructure:"WEB_SERVER_PORT"`
	OtelServiceName          string `mapstructure:"OTEL_SERVICE_NAME"`
	OtelExporterOtlpEndpoint string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	ServiceBUri              string `mapstructure:"SERVICE_B_URI"`
}

func LoadConfig(filePath string) *EnvConfig {
	var cfg *EnvConfig
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	log.Println("arquivo .env carregado")
	return cfg
}
