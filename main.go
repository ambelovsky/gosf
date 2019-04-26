package gosf

import (
	"log"
	"net/http"
	"runtime"
	"strconv"

	io "github.com/graarh/golang-socketio"
	transport "github.com/graarh/golang-socketio/transport"
)

// SupportedPlatforms is an array of OS platform names that this framework works well on
var SupportedPlatforms []string

// App is a global registry for application variables
var App map[string]interface{}

// Plugins is a global registry for framework plugins
var Plugins map[string]Plugin

// SocketIO Server
var server *io.Server

func init() {
	SupportedPlatforms = []string{"linux", "darwin", "windows"}
	if !ArrayContainsString(SupportedPlatforms, runtime.GOOS) {
		log.Panic("Unsupported platform.")
	}

	App = make(map[string]interface{})
	Plugins = make(map[string]Plugin)

	// SocketIO Server
	server = io.NewServer(transport.GetDefaultWebsocketTransport())
}

// RegisterPlugin registers a plugin for activation in the system
func RegisterPlugin(name string, plugin Plugin) {
	Plugins[name] = plugin
}

// Startup activates the framework and starts the server
func Startup(config map[string]interface{}) {
	port := 9999
	path := "/"

	if config["port"] != nil {
		port = config["port"].(int)
	}

	if config["path"] != nil {
		path = config["path"].(string)
	}

	// Activate configured plugins
	//TODO: pull from config file
	for _, plugin := range Plugins {
		plugin.Activate(&App)
	}

	// handle connected
	server.On(io.OnConnection, func(channel *io.Channel) {
		log.Println("Client connected")
	})

	// handle disconnected
	server.On(io.OnDisconnection, func(channel *io.Channel) {
		log.Println("Client disconnected")
	})

	// setup http server
	serveMux := http.NewServeMux()
	serveMux.Handle(path, server)
	log.Panic(http.ListenAndServe(":"+strconv.Itoa(port), serveMux))
}

// Shutdown cleanly terminates the framework and its plugins
func Shutdown() {
	// Deactivate configured plugins
	//TODO: pull from config file
	for _, plugin := range Plugins {
		plugin.Deactivate(&App)
	}
}

// Listen creates a listener on an endpoint
func Listen(endpoint string, function func(request *Request, clientMessage *Message)) {
	server.On(endpoint, func(channel *io.Channel, clientMessage *Message) {
		request := new(Request)
		request.Channel = channel
		request.Endpoint = endpoint

		for _, plugin := range Plugins {
			plugin.PreReceive(clientMessage)
		}

		function(request, clientMessage)

		for _, plugin := range Plugins {
			plugin.PostReceive(clientMessage)
		}
	})
}

// Respond sends a message back to the client
func (request Request) Respond(clientMessage *Message, serverMessage *Message) {
	channel := request.Channel

	for _, plugin := range Plugins {
		plugin.PreRespond(clientMessage, serverMessage)
	}

	if &clientMessage.ID != nil {
		serverMessage.ID = clientMessage.ID
	}

	channel.Emit(request.Endpoint, serverMessage)

	for _, plugin := range Plugins {
		plugin.PostRespond(clientMessage, serverMessage)
	}
}
