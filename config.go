package gosf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is a global registry for server configurations
var Config map[string]interface{}

// LoadConfig loads a JSON configuration file into the global Config map
func LoadConfig(name string, path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data map[string]interface{}
	json.Unmarshal([]byte(byteValue), &data)

	Config[name] = data
}

func init() {
	Config = make(map[string]interface{})
}
