package main

import (
	"github.com/eventbooking/lib/persistence/mongolayer"
)

func main() {

	db := mongolayer.InitMongoLayer("mongodb://localhost:27017/")
	defer db.Close()

}
