package gosf

import (
	"os"
	"reflect"
	"strings"
)

// AppSettings holds global settings for the application
type AppSettings struct {
	Env     map[string]string
	Config  map[string]interface{}
	Plugins map[string]reflect.Value
}

// App is a global registry for application variables
var App *AppSettings

func init() {
	App = new(AppSettings)
	App.Env = make(map[string]string)

	// Parse environmental variables
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		App.Env[pair[0]] = pair[1]
	}
}
