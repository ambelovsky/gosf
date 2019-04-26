package gosf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Config map[string]interface{}

func init() {
	Config = make(map[string]interface{})
}

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
