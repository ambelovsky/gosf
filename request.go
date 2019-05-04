package gosf

import (
	io "github.com/ambelovsky/gosf-socketio"
)

// Message - Standard message type for Socket communications
type Message struct {
	ID int `json:"id,omitempty"`

	Success bool   `json:"success"`
	Text    string `json:"text,omitempty"`

	Meta map[string]interface{} `json:"meta,omitempty"`
	Body map[string]interface{} `json:"body,omitempty"`
}

// WithoutMeta removes meta information before returning a copy of the message
func (m *Message) WithoutMeta() *Message {
	if m.Meta == nil {
		return m
	}

	var metaFreeMessage = new(Message)
	(*metaFreeMessage) = *m
	metaFreeMessage.Meta = nil
	return metaFreeMessage
}

// Request represents a single request over an active connection
type Request struct {
	Endpoint string
	Message  *Message
}

// Broadcast sends a message to connected clients joined to the same room
func Broadcast(room string, endpoint string, message *Message) {
	emit("before-broadcast", room, endpoint, message)
	if room != "" {
		ioServer.BroadcastTo(room, endpoint, message)
	} else {
		ioServer.BroadcastToAll(endpoint, message)
	}
	emit("after-broadcast", room, endpoint, message)
}

// Listen creates a listener on an endpoint
func Listen(endpoint string, callback func(client *Client, request *Request) *Message) {
	ioServer.On(endpoint, func(channel *io.Channel, clientMessage *Message) *Message {
		client := new(Client)
		client.channel = channel

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
	emit("before-response", client, &request, response)

	if response != nil {
		if &request.Message.ID != nil {
			response.ID = request.Message.ID
		}

		client.channel.Emit(request.Endpoint, response)
		return response
	}
	return nil
}
