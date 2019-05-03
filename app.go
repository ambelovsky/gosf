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

	log.Println("Hello world...")

	// Parse environmental variables
	for _, e := range os.Environ() {
		log.Println("found one")
		log.Println(e)

		pair := strings.Split(e, "=")
		App.Env[pair[0]] = pair[1]
	}
}
