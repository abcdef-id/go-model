package seeds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
)

// Seeder interface
type Seeder interface {
	Run(db *gorm.DB) error
}

// SeederList avaibale seeder
var SeederList = make(map[string]Seeder)

type jsonParser struct{}

func (jsonParser) ParseJSON(file string, output interface{}) {
	appPath, _ := os.Getwd()
	path := appPath + "/database/seeds/"
	raw, err := ioutil.ReadFile(path + file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(raw, &output)
}
