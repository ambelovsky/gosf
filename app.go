package gosf

// App is a global registry for application variables
var App map[string]interface{}

func init() {
	App = make(map[string]interface{})
}
