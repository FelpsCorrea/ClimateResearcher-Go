package configs

import "github.com/spf13/viper"

type conf struct {
	WeatherAPIBaseURL string `mapstructure:"WEATHER_API_BASE_URL"`
	WeatherAPIKey     string `mapstructure:"WEATHER_API_KEY"`
	ViaCepAPIBaseURL  string `mapstructure:"VIACEP_API_BASE_URL"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
