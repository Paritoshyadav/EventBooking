package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eventbooking/lib/configuration"
	"github.com/eventbooking/lib/persistence/dblayer"
	"github.com/eventbooking/rest"
)

func main() {

	conf_path := flag.String("conf", `lib\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*conf_path)
	fmt.Println("Connecting to database")
	dbhanlder, err, close := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	defer close()
	if err != nil {
		fmt.Println(err)
	}
	log.Fatal(rest.Server(config.RestfulEndpoint, dbhanlder))
}
