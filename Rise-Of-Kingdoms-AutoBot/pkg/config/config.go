package config

import "github.com/spf13/viper"

// Load config from file config with viper
type Configuration struct {
	ScreenSize    []int  `mapstructure:"screenSize"`
	Method        string `mapstructure:"method"`
	HaoiUser      string `mapstructure:"haoiUser"`
	HaoiRebate    string `mapstructure:"haoiRebate"`
	TwoCaptchaKey string `mapstructure:"twocaptchaKey"`
}

func LoadConfig() (*Configuration, error) {
	var config Configuration
	// Load the config file using viper
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
