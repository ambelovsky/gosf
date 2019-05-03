package gosf

type plugins map[string]struct {
	activation Plugin
}

func init() {
	App.Plugins = make(map[string]struct {
		activation Plugin
	})
}

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *AppSettings)
	Deactivate(app *AppSettings)
}

// RegisterPlugin registers a plugin for activation in the system
func RegisterPlugin(name string, plugin struct {
	activation Plugin
}) {
	App.Plugins[name] = plugin
}
