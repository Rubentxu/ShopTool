package main

import (
	prod "ShopTool/products" // import "github.com/Rubentxu/ShopTool"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/spf13/viper"
)

type productConfig struct {
	serverHost string
	serverPort string
	dbHost     string
	dbPort     string
}

func (c *productConfig) mongoAddr() string {
	return fmt.Sprintf("%s:%s", c.dbHost, c.dbPort)
}

func (c *productConfig) serverAddr() string {
	return fmt.Sprintf("%s:%s", c.serverHost, c.serverPort)
}

func config() (*productConfig, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("application")     // name of config file (without extension)
	viper.AddConfigPath("/etc/ShopTool/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.ShopTool") // call multiple times to add many search paths
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")  // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, fmt.Errorf("Fatal error config file: %s ", err)
	}
	println("config files")
	return &productConfig{
		serverHost: viper.GetString("microservices.products.server.host"),
		serverPort: viper.GetString("microservices.products.server.port"),
		dbHost:     viper.GetString("microservices.products.database.host"),
		dbPort:     viper.GetString("microservices.products.database.port"),
	}, nil

}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	config, err := config()
	if err != nil { // Handle errors reading the config file
		panic(err)
	}
	fmt.Printf("server address %s\n", config.serverAddr())
	fmt.Printf("mongo address %s\n", config.mongoAddr())
	h, _ := prod.NewHandler(config.mongoAddr())
	logger.Log(http.ListenAndServe(config.serverAddr(), h))

}
