// Copyright © 2017 Ignacio Coterillo <ignacio.coterillo@gmail.com>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "reflect"
	
    log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

type Instance map[string] interface{}

func initHTTPClient() *http.Client {
	config := viper.GetViper()
    tls_conf := config.Sub("tls")
    ca_file, err := ioutil.ReadFile(tls_conf.GetString("ca"))
    if err != nil {
        log.Fatal(err)
    }
    roots := x509.NewCertPool()
    ok := roots.AppendCertsFromPEM([]byte(ca_file))
    if !ok {
        log.Fatal("Failed to parse CA certificate")
    }
    
    // Initialize Transport and HTTP Client
    tlsConf := &tls.Config{RootCAs: roots}
    tr := &http.Transport{TLSClientConfig: tlsConf}
    return &http.Client{Transport: tr}
}

func GetInstance(instance string) Instance {

    log.Debug("Dumping " + instance)
    // Find and read the config file
	config := viper.GetViper()
	url := fmt.Sprintf("%s/%s/metadata", config.Get("api_instance_uri"), instance)
    log.Debug("API URL:" + url)

    client := initHTTPClient()
    res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

    body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
    var resp = make(map[string] []Instance, 0)
	if err != nil {
		log.Fatal(err)
	} else {
        log.Debug(reflect.TypeOf(resp))
        log.Debug(string(body))
        json.Unmarshal(body, &resp)
    }

    // API Response is an array of JSON Objects
    return resp["response"][0]

}

