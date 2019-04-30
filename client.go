package gosf

import io "github.com/ambelovsky/gosf-socketio"

// Client represents a single connected client
type Client struct {
	channel *io.Channel
	Rooms   []string
}

// Join joins a user to a broadcast room
func (c *Client) Join(room string) {
	c.channel.Join(room)

	foundRoom := false
	for i := range c.Rooms {
		if c.Rooms[i] == room {
			foundRoom = true
			break
		}
	}

	if !foundRoom {
		c.Rooms = append(c.Rooms, room)
	}
}

// Leave removes a user from a broadcast room
func (c *Client) Leave(room string) {
	c.channel.Leave(room)

	for i := range c.Rooms {
		if c.Rooms[i] == room {
			c.Rooms = append(c.Rooms[:i], c.Rooms[i+1:]...)
			break
		}
	}
}

// LeaveAll removes a user from all broadcast rooms they are currently joined to
func (c *Client) LeaveAll() {
	for _, v := range c.Rooms {
		c.channel.Leave(v)
	}

	c.Rooms = make([]string, 0)
}

// Disconnect forces a client to be disconnected from the server
func (c *Client) Disconnect() {
	c.channel.Close()
}

// Broadcast sends a message to connected clients joined to the same room
// with the exception of the "client"
func (c *Client) Broadcast(room string, endpoint string, message *Message) {
	emit("before-client-broadcast", c, room, endpoint, message)
	c.channel.BroadcastTo(room, endpoint, message)
	emit("after-client-broadcast", c, room, endpoint, message)
}
