package cmd

import (
	"encoding/json"
	"fmt"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"

	log "github.com/Sirupsen/logrus"
	"github.com/icot/dbod/api"
	"github.com/icot/dbod/config"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.Load()
}

func TestDump(t *testing.T) {
	ass := assert.New(t)

	client := api.GetClient()
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
	var buf = make(map[string][]api.Instance, 0)
	json.Unmarshal([]byte(target_response), &buf)

	target_instance := buf["response"][0]
	log.Debug("Target instance: ", target_instance)

	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, target_response))
	metadata := api.GetInstance(instance)
	ass.Equal(metadata, target_instance, "Instance body should match")

}
