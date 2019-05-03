package gosf

type plugins map[string]Plugin

func init() {
	App.Plugins = make(map[string]Plugin)
}

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *AppSettings)
	Deactivate(app *AppSettings)
}

// RegisterPlugin registers a plugin for activation in the system
func RegisterPlugin(name string, plugin Plugin) {
	App.Plugins[name] = plugin
}
