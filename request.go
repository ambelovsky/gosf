package gosf

import (
	io "github.com/ambelovsky/gosf-socketio"
)

// Message - Standard message type for Socket communications
type Message struct {
	ID      int  `json:"id,omitempty"`
	Success bool `json:"success"`

	Text string      `json:"text,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

// Request represents a single request over an active connection
type Request struct {
	Channel  *io.Channel
	Endpoint string
	Message  *Message
}

// Listen creates a listener on an endpoint
func Listen(endpoint string, callback func(request *Request) *Message) {
	server.On(endpoint, func(channel *io.Channel, clientMessage *Message) *Message {
		request := new(Request)
		request.Channel = channel
		request.Endpoint = endpoint
		request.Message = clientMessage

		for _, plugin := range Plugins {
			plugin.PreRequest(request)
		}

		response := callback(request)

		for _, plugin := range Plugins {
			plugin.PostRequest(request, response)
		}

		return request.respond(response)
	})
}

// Respond sends a message back to the client
func (request Request) respond(serverMessage *Message) *Message {
	channel := request.Channel

	for _, plugin := range Plugins {
		plugin.PreResponse(&request, serverMessage)
	}

	if &request.Message.ID != nil {
		serverMessage.ID = request.Message.ID
	}

	channel.Emit(request.Endpoint, serverMessage)

	for _, plugin := range Plugins {
		plugin.PostResponse(&request, serverMessage)
	}

	return serverMessage
}
