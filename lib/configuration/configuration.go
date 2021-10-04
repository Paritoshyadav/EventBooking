package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/eventbooking/lib/persistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://localhost:27017/"
	RestfulEPDefault    = ":8080"
)

type ServiceConfig struct {
	Databasetype    dblayer.DBTYPE `json:"databasetype"`
	DBConnection    string         `json:"dbconnection"`
	RestfulEndpoint string         `json:"restfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	Sc := ServiceConfig{
		Databasetype:    DBTypeDefault,
		DBConnection:    DBConnectionDefault,
		RestfulEndpoint: RestfulEPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return Sc, err
	}
	err = json.NewEncoder(file).Encode(&Sc)

	return Sc, err

}
