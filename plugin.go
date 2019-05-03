package gosf

import "reflect"

type pluginRegistration struct {
	active     bool
	plugin     Plugin
	appMethods reflect.Value
}

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *AppSettings)
	Deactivate(app *AppSettings)
}

var plugins map[string]*pluginRegistration

func init() {
	plugins = make(map[string]*pluginRegistration)
	App.Plugins = make(map[string]reflect.Value)
}

// RegisterPlugin is used by the plugin to register itself for later activation
func RegisterPlugin(name string, plugin Plugin, appMethods interface{}) {
	plugins[name] = new(pluginRegistration)
	plugins[name].active = false
	plugins[name].plugin = plugin
	plugins[name].appMethods = reflect.ValueOf(appMethods)
}

// LoadPlugin activates a registered plugin
func LoadPlugin(systemName string, referenceName string) {
	if referenceName == "" {
		referenceName = systemName
	}

	plugins[systemName].active = true
	App.Plugins[referenceName] = reflect.ValueOf(plugins[systemName].appMethods)
}
