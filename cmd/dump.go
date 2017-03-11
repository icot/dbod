// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
    "encoding/json"
	"fmt"

    log "github.com/Sirupsen/logrus"
    "github.com/icot/dbod/api"
	"github.com/spf13/cobra"
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
	log.Debug("dump called")
	if len(args) != 1 {
		log.Fatal("Error: please run $ dbod " + dumpCmd.Use )
	}
	instance := args[0]
    metadata := api.GetInstance(instance)
    str, _ := json.MarshalIndent(metadata, "", "  ")
    fmt.Println(string(str))
}


