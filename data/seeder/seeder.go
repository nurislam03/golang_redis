package seeder

import (
	"log"
	"reflect"

	"github.com/kamva/mgm/v3"

	"github.com/nurislam03/template/pkg/dbconn"
)

type Seed struct{}

func Execute(seedMethodNames []string) {

	err := dbconn.Connect()
	if err != nil {
		log.Fatal("mongodb connect", err)
	}

	_, client, _, _ := mgm.DefaultConfigs()
	defer client.Disconnect(mgm.Ctx())

	s := Seed{}
	seedType := reflect.TypeOf(s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succeed")
}
