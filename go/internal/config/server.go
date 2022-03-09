package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Server è una struttura con le impostazioni del server.
type Server struct {
	Host string `yaml:"host" json:"host,omitempty"`
	Port string `yaml:"port" json:"port,omitempty"`
	Ssl  SSLConfig `yaml:"ssl" json:"ssl,omitempty"`
	Limiter int `yaml:"limiter" json:"limiter,omitempty"`
	Gui      string `yaml:"gui" json:"gui,omitempty"`
	//Test     string `yaml:"test" json:"test,omitempty"`
	//Template string `yaml:"template" json:"template,omitempty"`
}


//SSLConfig è una struttura con le impostazioni dei certificati SSL.
type SSLConfig struct {
	Certificate string `yaml:"certificate" json:"certificate,omitempty"`
	Key         string `yaml:"key" json:"key,omitempty"`
}

//CheckServer controlla la validità di tutte le impostazioni relative al server.
func (srv *Server) CheckServer() error {

	//Check host validity
	if err := srv.CheckHost(); err != nil {
		return err
	}

	//Check port validity and disponibility
	if err := srv.CheckPort(); err != nil {
		return err
	}

	//Check Limiter validity
	if err := srv.CheckLimiter(); err != nil {
		return err
	}

	/*
	//TODO: Abilitare se vuoi fare un server HTTPS
	//Check SSL config
	if err := srv.Ssl.CheckSSL(); err != nil {
		return err
	}
	*/

	/*
	//Check template path
	if err := srv.CheckTemplate(); err != nil {
		return err
	}
	*/

	//Check gui path
	if err := srv.CheckGui(); err != nil {
		return err
	}

	/*
	//Check test path
	if err := srv.CheckTest(); err != nil {
		return err
	}
	 */

	return nil
}

//CheckHost controlla l'impostazione host del server.
func (srv *Server) CheckHost() error {
	if srv.Host == "" {
		return errors.New("Server host not valid.")
	}
	return nil
}

//CheckPort controlla l'impostazione porta del server.
func (srv *Server) CheckPort() error {
	portInt, err := strconv.Atoi(strings.Trim(srv.Port, ":"))
	if err != nil {
		return errors.New("Invalid Port: Conversion of port to int not valid")
	}

	if portInt < 1024 || portInt > 49151 {
		return errors.New("Port not valid. [1024-49151]")
	}
	return nil
}

/*
//CheckTemplate controlla la path template del server.
func (srv *ServerConfig) CheckTemplate() error {
	if _, err := os.Stat(srv.Template); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf(`Template path "%s", not exist. Modify config file "server.yml!"`, srv.Template))
	}
	return nil
}
 */

//CheckGui controlla la path gui del server.
func (srv *Server) CheckGui() error {
	if _, err := os.Stat(srv.Gui); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf(`Gui path "%s", not exist. Modify config file "server.yml!"`, srv.Gui))
	}
	return nil
}

/*
//CheckTest controlla la path test del server
func (srv *ServerConfig) CheckTest() error {
	if _, err := os.Stat(srv.Test); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf(`Test path "%s", not exist. Modify config file "server.yml!"`, srv.Test))
	}
	return nil
}
*/


//CheckSSL controlla tutte le impostazioni SSL del server.
func (ssl *SSLConfig) CheckSSL() error {

	//Check key SSL
	if _, err := os.Stat(ssl.Key); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf(`"%s": Not exist! Modify config file "server.yml!"`, ssl.Key))
	}

	//Check certificate SSL
	if _, err := os.Stat(ssl.Certificate); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf(`"%s": Not exist! Modify config file "server.yml!"`, ssl.Certificate))
	}

	return nil
}

//CheckLimiter controlla tutte le impostazioni del limiter del server.
func (srv *Server) CheckLimiter() error {

	if srv.Limiter < 1 {
		return errors.New("Limiter not valid.")
	}

	return nil
}
