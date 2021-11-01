package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

type ApplicationEnvironment struct {
	TemplateLocation string
}

func ConfigSetup(configName, configPath string) {
	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
}

func ValidateVariablesAreSet(variables []string) {
	for i := range variables {
		if !viper.IsSet(variables[i]) {
			log.Fatalf("%s variable was not set!\nAborting application start!", variables[i])
		}
	}
}

func GetConfig() ApplicationEnvironment {

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file in infrastructure: %s", err)
	}
	ValidateVariablesAreSet([]string{})

	return ApplicationEnvironment{
		TemplateLocation: viper.GetString("TemplateLocation"),
	}
}
