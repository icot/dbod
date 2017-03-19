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

var host_metadata = []byte(`[{"username":"icoteril","version":"5.6.17","db_type":"MYSQL","basedir":"/usr/local/mysql/mysql-5.6.17","class":"TEST","port":"5500","datadir":"/ORA/dbs03/PINOCHO/mysql","db_name":"pinocho","hosts":["db-50019"],"logdir":"/ORA/dbs02/PINOCHO/mysql","volumes":[{"group":"mysql","file_mode":"0755","server":"dbnasr0010-priv","instance_id":20,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"mysql","mounting_path":"/ORA/dbs02/PINOCHO","id":189},{"group":"mysql","file_mode":"0755","server":"dbnasr0003-priv","instance_id":20,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"mysql","mounting_path":"/ORA/dbs03/PINOCHO","id":190}],"attributes":{"notifications":"false","port":"5500","buffer_pool_size":"1G"},"id":20,"bindir":"/usr/local/mysql/mysql-5.6.17/bin","socket":"/var/lib/mysql/mysql.sock.pinocho.5500"},{"username":"adumitru","version":"5.6.17","db_type":"MYSQL","basedir":"/usr/local/mysql/mysql-5.6.17","class":"TEST","port":"5505","datadir":"/ORA/dbs03/ADUMITRU/mysql","db_name":"adumitru","hosts":["db-50019"],"logdir":"/ORA/dbs02/ADUMITRU/mysql","volumes":[{"group":"mysql","file_mode":"0755","server":"dbnasr0014-priv","instance_id":256,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"mysql","mounting_path":"/ORA/dbs02/ADUMITRU","id":602},{"group":"mysql","file_mode":"0755","server":"dbnasr0004-priv","instance_id":256,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"mysql","mounting_path":"/ORA/dbs03/ADUMITRU","id":603}],"attributes":{"notifications":"true","port":"5505","buffer_pool_size":"1G"},"id":256,"bindir":"/usr/local/mysql/mysql-5.6.17/bin","socket":"/var/lib/mysql/mysql.sock.adumitru.5505"},{"username":"icoteril","version":"5.6.17","db_type":"MYSQL","basedir":"/usr/local/mysql/mysql-5.6.17","class":"TEST","port":"5501","datadir":"/ORA/dbs03/PINOCHO2/mysql","db_name":"pinocho2","hosts":["db-50019"],"logdir":"/ORA/dbs02/PINOCHO2/mysql","volumes":[{"group":"mysql","file_mode":"0755","server":"dbnasr0015-priv","instance_id":277,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"mysql","mounting_path":"/ORA/dbs02/PINOCHO2","id":261},{"group":"mysql","file_mode":"0755","server":"dbnasr0004-priv","instance_id":277,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"mysql","mounting_path":"/ORA/dbs03/PINOCHO2","id":262}],"attributes":{"notifications":"true","port":"5501","buffer_pool_size":"1G"},"id":277,"bindir":"/usr/local/mysql/mysql-5.6.17/bin","socket":"/var/lib/mysql/mysql.sock.pinocho2.5501"},{"username":"jocorder","version":"9.4.5","db_type":"PG","basedir":"/usr/local/pgsql/pgsql-9.4.5","class":"TEST","port":"6601","datadir":"/ORA/dbs03/PGTEST/data","db_name":"pgtest","hosts":["db-50019"],"logdir":"/ORA/dbs02/PGTEST/pg_xlog","volumes":[{"group":"postgres","file_mode":"0755","server":"dbnasr0015-priv","instance_id":297,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"postgres","mounting_path":"/ORA/dbs02/PGTEST","id":299},{"group":"postgres","file_mode":"0755","server":"dbnasr0004-priv","instance_id":297,"mount_options":"rw,bg,hard,nointr,tcp,vers=3,noatime,timeo=600,rsize=65536,wsize=65536","owner":"postgres","mounting_path":"/ORA/dbs03/PGTEST","id":300}],"attributes":{"notifications":"true","port":"6601","shared_buffers":"1G"},"id":297,"bindir":"/usr/local/pgsql/pgsql-9.4.5/bin","socket":"/var/lib/pgsql/"}]`)
