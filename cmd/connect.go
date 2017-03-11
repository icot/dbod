// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
    "os"
    "os/exec"
    "strings"
    "syscall"
	
    log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
    "github.com/icot/dbod/api"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect <instance>",
	Short: "Connect to an instance using an external tool",
}

func init() {
	RootCmd.AddCommand(connectCmd)
	connectCmd.Run = connect

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func connect (cmd *cobra.Command, args []string) {
	
    log.Debug(" connect called")
	if len(args) != 1 {
		log.Fatal("Error: please run $ dbod " + dumpCmd.Use )
	}
	instance := args[0]

    // Read CLI commands and arguments from configuration
	config := viper.GetViper()
	url := fmt.Sprintf("%s/%s/metadata", config.Get("api_instance_uri"), instance)
    log.Debug("API URL:" + url)

    // Fetch instance metadata
    metadata := api.GetInstance(instance)
    
    db_type := "mysql"
    fmt.Println(metadata["time"])

    // Load 
    cli := config.Sub("cli")
    cmd_line := cli.GetStringMapString(db_type)
    log.Debug("Client: ", cmd_line["client"])
    cmd_args := strings.Fields(fmt.Sprintf(cmd_line["args"], instance, "5500"))
    log.Debug("Cmd Line: ", cmd_args)
    // Look for binary
    binary, lookErr := exec.LookPath(cmd_line["client"])
    if lookErr != nil {
        log.Fatal(lookErr)
    }
    env := os.Environ()
    // Execute client
    execErr := syscall.Exec(binary, cmd_args, env)
    if execErr != nil {
        log.Fatal(execErr)
    }
}
