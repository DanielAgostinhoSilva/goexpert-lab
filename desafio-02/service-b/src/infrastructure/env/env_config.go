package env

import (
	"github.com/spf13/viper"
	"log"
)

type EnvConfig struct {
	WebServerPort            string `mapstructure:"WEB_SERVER_PORT"`
	WeatherApiKey            string `mapstructure:"WEATHER_API_KEY"`
	ViaCepApiUri             string `mapstructure:"VIA_CEP_URI"`
	WeatherApiUri            string `mapstructure:"WEATHER_URI"`
	OtelServiceName          string `mapstructure:"OTEL_SERVICE_NAME"`
	OtelExporterOtlpEndpoint string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"`
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
