package gosf

// Plugins is a global registry for framework plugins
var Plugins map[string]Plugin

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *map[string]interface{}, config *map[string]interface{})
	Deactivate(app *map[string]interface{}, config *map[string]interface{})
}

// RegisterPlugin registers a plugin for activation in the system
func RegisterPlugin(name string, plugin Plugin) {
	Plugins[name] = plugin
}

func init() {
	Plugins = make(map[string]Plugin)
}
