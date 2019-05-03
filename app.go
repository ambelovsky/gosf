package gosf

import (
	"log"
	"os"
	"strings"
)

// AppSettings holds global settings for the application
type AppSettings struct {
	Env    map[string]string
	Config map[string]interface{}
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

	for k, v := range App.Env {
		log.Println("Registered env var name " + k + " with value " + v)
	}
}
