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
	Endpoint string
	Message  *Message
}

// Broadcast sends a message to connected clients joined to the same room
func Broadcast(room string, endpoint string, message *Message) {
	if room != "" {
		ioServer.BroadcastTo(room, endpoint, message)
	} else {
		ioServer.BroadcastToAll(endpoint, message)
	}
}

// Listen creates a listener on an endpoint
func Listen(endpoint string, callback func(client *Client, request *Request) *Message) {
	ioServer.On(endpoint, func(channel *io.Channel, clientMessage *Message) *Message {
		client := new(Client)
		client.Channel = channel

		request := new(Request)
		request.Endpoint = endpoint
		request.Message = clientMessage

		emit("before-request", client, request)

		response := callback(client, request)

		emit("after-request", client, request, response)

		defer emit("after-response", client, request, response)

		return request.respond(client, response)
	})
}

// Respond sends a message back to the client
func (request Request) respond(client *Client, response *Message) *Message {
	if &request.Message.ID != nil {
		response.ID = request.Message.ID
	}

	emit("before-response", client, &request, response)

	if response != nil {
		client.Channel.Emit(request.Endpoint, response)
		return response
	} else {
		return nil
	}
}
