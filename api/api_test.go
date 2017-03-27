package api

import (
	"encoding/json"
	"fmt"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/icot/dbod/config"

	assert "github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
)

var cfgFile string

func init() {
	config.Load()
}

func TestGetInstanceSuccess(t *testing.T) {
	ass := assert.New(t)

	client := GetClient()
	log.Debug(fmt.Sprintf("Initialized API Client[%p]", client))
	log.Debug(fmt.Sprintf("Activating httpmock"))
	httpmock.ActivateNonDefault(client)
	defer httpmock.DeactivateAndReset()

	instance := "myt"
	uri := "https://api-server/api/v1/instance"
	url := fmt.Sprintf("%s/%s/metadata", uri, instance)

	// Succesful request
	var target_response = `{"response":[{"id": 1, "db_name": "myt"}]}`
	log.Debug("Target response: ", target_response)
	var buf = make(map[string][]Instance, 0)
	json.Unmarshal([]byte(target_response), &buf)

	target_instance := buf["response"][0]
	log.Debug("Target instance: ", target_instance)

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, target_response))
	metadata := GetInstance(instance)
	ass.Equal(metadata, target_instance, "Instance body should match")

	// Unsuccesful request
	target_response = ``
	log.Debug("Target response: ", target_response)
	json.Unmarshal([]byte(target_response), &buf)
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(404, target_response))
	metadata = GetInstance(instance)
	ass.Equal(metadata, Instance(nil), "Instance body should be empty")
}
