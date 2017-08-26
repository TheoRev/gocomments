package configuration

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	// Driver de postgresql
	_ "github.com/lib/pq"
)

// Configuration estructura de Configuration
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfiguration obtiene el json de configuracion con la db
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func GetConnection() *gorm.DB {
	c := GetConfiguration()
}
