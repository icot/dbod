package config

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/spf13/viper"
)

var cfgFile string

func init() {
}

// initConfig reads in config file and ENV variables if set.
func Load() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName(".dbodrc") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AddConfigPath(".")       // optionally look for config in the working directory
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal((fmt.Errorf("Fatal error config file: %s \n", err)))
	}

	log.SetLevel(log.DebugLevel)
}
