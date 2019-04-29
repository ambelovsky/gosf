package gosf

// Plugins is a global registry for framework plugins
var Plugins map[string]Plugin

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *map[string]interface{}, config *map[string]interface{})
	Deactivate(app *map[string]interface{}, config *map[string]interface{})

	Connect(request *Request)
	Disconnect(request *Request)

	PreRequest(request *Request)
	PostRequest(request *Request, response *Message)
	PreResponse(request *Request, response *Message)
	PostResponse(request *Request, response *Message)
}

// RegisterPlugin registers a plugin for activation in the system
func RegisterPlugin(name string, plugin Plugin) {
	Plugins[name] = plugin
}

func init() {
	Plugins = make(map[string]Plugin)
}
