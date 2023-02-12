package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config is the server configuration structure.
// all fields will be filled with environment variables.
type Config struct {
	ServerHost    string `mapstructure:"SERVER_HOST"` // address that server will listening on
	MongoUser     string `mapstructure:"MONGO_USER"` // mongo db username
	MongoPassword string `mapstructure:"MONGO_PASSWORD"` // mongo db password
	MongoHost     string `mapstructure:"MONGO_HOST"` // host that mongo db listening on
	MongoPort     string `mapstructure:"MONGO_PORT"` // port that mongo db listening on
}

// initialize will read environment variables and save them in config structure fields
func (config *Config) initialize() {
	// read environment variables
	// config.ServerHost = os.Getenv("server_host")
	// config.MongoUser = os.Getenv("mongo_user")
	// config.MongoPassword = os.Getenv("mongo_password")
	// config.MongoHost = os.Getenv("mongo_host")
	// config.MongoPort = os.Getenv("mongo_port")

	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.SetConfigName("mongo")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file: %s", err)
	}

	err = viper.Unmarshal(&config)
}

// MongoURI will generate mongo db connect uri
func (config *Config) MongoURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.MongoUser,
		config.MongoPassword,
		config.MongoHost,
		config.MongoPort,
	)
}

// NewConfig will create and initialize config struct
func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
