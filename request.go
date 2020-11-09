package gosf

import (
	io "github.com/ambelovsky/gosf-socketio"
)

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
		if &request.Message.GUID != nil {
			response.GUID = request.Message.GUID
		}
		if &request.Message.UUID != nil {
			response.UUID = request.Message.UUID
		}

		client.channel.Emit(request.Endpoint, response)
		return response
	}
	return nil
}
