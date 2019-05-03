package gosf

import (
	"log"
	"net/http"
	"reflect"
	"strconv"

	io "github.com/ambelovsky/gosf-socketio"
	transport "github.com/ambelovsky/gosf-socketio/transport"
)

// SocketIO Server
var ioServer *io.Server

func init() {
	// Transport configuration
	transport := transport.GetDefaultWebsocketTransport()
	transport.UnsecureTLS = true

	// SocketIO Server
	ioServer = io.NewServer(transport)
}

// Startup activates the framework and starts the server
func Startup(config map[string]interface{}) {
	secure := false
	port := 9999
	path := "/"

	/*** CONFIG ***/

	if config["secure"] != nil {
		secure = config["secure"].(bool)
	}

	if config["port"] != nil {
		if reflect.TypeOf(config["port"]).String() == "float64" {
			port = int(config["port"].(float64))
		} else {
			port = config["port"].(int)
		}
	}

	if config["path"] != nil {
		path = config["path"].(string)
	}

	address := ":" + strconv.Itoa(port)

	if config["host"] != nil {
		address = config["host"].(string) + address
	}

	/*** END CONFIG ***/

	// Activate configured plugins
	for _, plugin := range App.Plugins {
		plugin.activation.Activate(App)
	}

	// Start connection handlers
	startConnectionHandlers()

	// Setup http server
	startHTTPServer(config, secure, address, port, path)
}

// Shutdown cleanly terminates the framework and its plugins
func Shutdown() {
	// Deactivate configured plugins
	for _, plugin := range App.Plugins {
		plugin.activation.Deactivate(App)
	}
}

func startConnectionHandlers() {
	// Handle connected
	ioServer.On(io.OnConnection, func(channel *io.Channel) {
		client := new(Client)
		client.channel = channel

		request := new(Request)
		request.Endpoint = "connect"

		emit("connect", client, request)
	})

	// Handle disconnected
	ioServer.On(io.OnDisconnection, func(channel *io.Channel) {
		client := new(Client)
		client.channel = channel

		request := new(Request)
		request.Endpoint = "disconnect"

		emit("disconnect", client, request)
	})
}

func startHTTPServer(config map[string]interface{}, secure bool, address string, port int, path string) {
	serveMux := http.NewServeMux()
	serveMux.Handle(path, ioServer)

	if !secure || config["ssl-cert"] == nil || config["ssl-key"] == nil {
		log.Panic(http.ListenAndServe(address, serveMux))
	} else {
		log.Panic(http.ListenAndServeTLS(address, config["ssl-cert"].(string), config["ssl-key"].(string), serveMux))
	}
}
