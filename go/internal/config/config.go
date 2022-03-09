package config

import (
	"go.uber.org/config"
)

//Config è una struttura con le impostazioni del software.
type Config struct {
	Server Server `yaml:"server" json:"server,omitempty"`
	//Logger LoggerConfig `yaml:"logger" json:"logger,omitempty"`
	Database Database `yaml:"database" json:"database,omitempty"`
	//Password PasswordConfig `yaml:"password" json:"password,omitempty"`
	//Mail MailConfig `yaml:"mail" json:"mail,omitempty"`
	//Template TemplateConfig `yaml:"template" json:"template,omitempty"`
	//Jwt JWTConfig `yaml:"jwt" json:"jwt,omitempty"`
}

//NewConfig è un costruttore dell'oggetto config.
func NewConfig(pathConfig string) (*Config, error) {

	provider, err := config.NewYAML(
		config.File(pathConfig+"server.yml"),
		//config.File(pathConfig+"logger.yml"),
		config.File(pathConfig+"database.yml"),
		//config.File(pathConfig+"password.yml"),
		//config.File(pathConfig+"mail.yml"),
		//config.File(pathConfig+"template.yml"),
		//config.File(pathConfig+"jwt.yml"),
	)

	if err != nil {
		return nil, err
	}

	var c Config
	if err := provider.Get("").Populate(&c); err != nil {
		panic(err) // handle error
	}

	return &c, nil
}

//CheckConfig controlla la validità di tutte le configurazioni.
func (cfg *Config) CheckConfig() error {
	//Check server settings
	if err := cfg.Server.CheckServer(); err != nil {
		return err
	}

	/*
	//Check logger settings
	if err := cfg.Logger.CheckLogger(); err != nil {
		return err
	}
	*/

	//Check not relational databases settings
	if err := cfg.Database.CheckDatabase(); err != nil {
			return err
		}

	/*
	//Check password settings
	if err := cfg.Password.CheckPassword(); err != nil {
		return err
	}

	//Check template settings
	if err := cfg.Template.CheckTemplate(); err != nil {
		return err
	}
	*/

	return nil
}