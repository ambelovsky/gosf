package gosf

// System plugin activators
var plugins map[string]*Plugin

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *AppSettings)
	Deactivate(app *AppSettings)
}

func init() {
	plugins = make(map[string]*Plugin)
	App.Plugins = make(map[string]*interface{})
}

// RegisterPlugin registers a plugin for activation in the system
func RegisterPlugin(name string, plugin *Plugin, appMethods *interface{}) {
	plugins[name] = plugin
	App.Plugins[name] = appMethods
}
