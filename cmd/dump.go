// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump <instance>",
	Short: "Dumps an instance metadata",
}

func init() {
	RootCmd.AddCommand(dumpCmd)
	dumpCmd.Run = dump
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//dumpCmd.PersistentFlags().String("instance", "", "database instance name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func dump (cmd *cobra.Command, args []string) {
	fmt.Println("dump called")
	config := viper.GetViper() // Find and read the config file
	if len(args) != 1 {
		log.Fatal("Error: please run $ dbod " + dumpCmd.Use )
	}
	instance := args[0]
	url := fmt.Sprintf("%s/%s/metadata", config.Get("api_instance_uri"), instance)
    log.Debug("API URL:" + url)
}
