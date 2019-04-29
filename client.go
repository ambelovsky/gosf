package gosf

import io "github.com/ambelovsky/gosf-socketio"

type Client struct {
	Channel *io.Channel
}

// Join joins a user to a broadcast room
func (c *Client) Join(room string) {
	c.Channel.Join(room)
}

// Leave removes a user from a broadcast room
func (c *Client) Leave(room string) {
	c.Channel.Leave(room)
}

// Broadcast sends a message to connected clients joined to the same room
// with the exception of the "client"
func (c *Client) Broadcast(room string, endpoint string, message *Message) {
	c.Channel.BroadcastTo(room, endpoint, message)
}
