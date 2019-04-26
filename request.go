package gosf

import io "github.com/ambelovsky/gosf-socketio"

// Message - Standard message type for Socket communications
type Message struct {
	ID      int  `json:"id,omitempty"`
	Success bool `json:"success"`

	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Request represents a single request over an active connection
type Request struct {
	Channel  *io.Channel
	Endpoint string
}

// Listen creates a listener on an endpoint
func Listen(endpoint string, function func(request *Request, clientMessage *Message)) {
	server.On(endpoint, func(channel *io.Channel, clientMessage *Message) {
		request := new(Request)
		request.Channel = channel
		request.Endpoint = endpoint

		for _, plugin := range Plugins {
			plugin.PreReceive(request, clientMessage)
		}

		function(request, clientMessage)

		for _, plugin := range Plugins {
			plugin.PostReceive(request, clientMessage)
		}
	})
}

// Respond sends a message back to the client
func (request Request) Respond(clientMessage *Message, serverMessage *Message) {
	channel := request.Channel

	for _, plugin := range Plugins {
		plugin.PreRespond(&request, clientMessage, serverMessage)
	}

	if &clientMessage.ID != nil {
		serverMessage.ID = clientMessage.ID
	}

	channel.Emit(request.Endpoint, serverMessage)

	for _, plugin := range Plugins {
		plugin.PostRespond(&request, clientMessage, serverMessage)
	}
}
