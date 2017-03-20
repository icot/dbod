package api

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"testing"

	"github.com/spf13/viper"
	"gopkg.in/jarcoal/httpmock.v1"
)

var cfgFile string
var DebugFlag bool

func init() {
	initConfig()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

func TestGetInstance(t *testing.T) {
	client := GetClient()
	log.Debug(fmt.Sprintf("Initialized API Client[%p]", client))
	log.Debug(fmt.Sprintf("Activating httpmock"))
	httpmock.ActivateNonDefault(client)
	defer httpmock.DeactivateAndReset()

	instance := "myt"
	uri := "https://api-server/api/v1/instance"
	url := fmt.Sprintf("%s/%s/metadata", uri, instance)

	log.Info("Mocking " + url)
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, `{"response":[{"id": 1, "name": "My Great Article"}]}`))

	metadata := GetInstance(instance)
	fmt.Println(metadata)

}
