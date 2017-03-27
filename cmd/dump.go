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
	"github.com/hokaccha/go-prettyjson"
	"github.com/icot/dbod/api"
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump <instance>",
	Short: "dumps an instance's metadata",
}

func init() {
	RootCmd.AddCommand(dumpCmd)
	dumpCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Error: please run $ dbod " + dumpCmd.Use)
		}
		dump(args[0])
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//dumpCmd.PersistentFlags().String("instance", "", "database instance name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// dump implements the main command functionality as an externally callable
// function to make it easily testable
func dump(instance_name string) {

	metadata := api.GetInstance(instance_name)
	if metadata != nil {
		str, _ := prettyjson.Marshal(metadata)
		fmt.Println(string(str))
	}
}
