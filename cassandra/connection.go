package cassandra

import (
	"crypto/tls"
	"crypto/x509"
	//"errors"
	"log"
	"time"
	"github.com/gocql/gocql"
	//"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func CheckUserPassword(useSSL bool,username string,password string,port int,timeout int,protocolVersion int,rawHosts []interface{},rootCA string, minTLSVersion (string) ) (string) {
	log.Printf("Checking Password")
	log.Printf("Using port %d", port)
	log.Printf("Using use_ssl %v", useSSL)
	log.Printf("Using username %s", username)

	hosts := make([]string, len(rawHosts))

	for _, value := range rawHosts {
		hosts = append(hosts, value.(string))

		log.Printf("Using host %v", value.(string))
	}

	cluster := gocql.NewCluster()

	cluster.Hosts = hosts

	cluster.Port = port

	cluster.Authenticator = &gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}

	cluster.ConnectTimeout = time.Millisecond * time.Duration(timeout)

	cluster.Timeout = time.Minute * time.Duration(1)

	cluster.CQLVersion = "3.0.0"

	cluster.Keyspace = "system"

	cluster.ProtoVersion = protocolVersion

	cluster.HostFilter = gocql.WhiteListHostFilter(hosts...)

	cluster.DisableInitialHostLookup = true

	log.Printf("Chegou aqui 1")

	if useSSL {

		tlsConfig := &tls.Config{
			MinVersion: allowedTlsProtocols[minTLSVersion],
		}

		if rootCA != "" {
			caPool := x509.NewCertPool()
			//ok := caPool.AppendCertsFromPEM([]byte(rootCA))

			//if !ok {
			//	return nil, errors.New("Unable to load rootCA")
			//}

			tlsConfig.RootCAs = caPool
		}

		cluster.SslOpts = &gocql.SslOptions{
			Config: tlsConfig,
		}
	}

	sessionCheck, sessionCheckCreateError := cluster.CreateSession()

	//elapsed := time.Since(start)

	//log.Printf("Getting a session took %s", elapsed)
	log.Printf("Chegou aqui 2")
	defer sessionCheck.Close()

	if sessionCheckCreateError != nil {
		return "change_me"
	} else {
		return password
	}
}
