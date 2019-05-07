package gosf

import (
	"encoding/json"
	"time"

	io "github.com/ambelovsky/gosf-socketio"
	transport "github.com/ambelovsky/gosf-socketio/transport"
)

func init() {
	App.Microservices = make(map[string]*Microservice)
}

// Microservice represents a microservice connection
type Microservice struct {
	host      string
	port      int
	secure    bool
	transport *transport.WebsocketTransport
	client    *io.Client
}

func (m *Microservice) configure(host string, port int, secure bool) *Microservice {
	m.host = host
	m.port = port
	m.secure = secure
	m.transport = transport.GetDefaultWebsocketTransport()

	return m
}

// Connected tells whether or not this microservice's connection is still active and alive
func (m *Microservice) Connected() bool {
	if m.client == nil {
		return false
	}
	return m.client.IsAlive()
}

// Connect manually connects or reconnects to the microservice
func (m *Microservice) Connect() (*Microservice, error) {
	if m.Connected() {
		m.Disconnect()
	}

	var err error
	url := io.GetUrl(m.host, m.port, m.secure)
	m.client, err = io.Dial(url, m.transport)
	return m, err
}

// Disconnect manually terminates the connection to the microservice
func (m *Microservice) Disconnect() {
	if m.Connected() == false {
		return
	}
	m.client.Close()
}

// Listen registers and event handler for a microservice endpoint
func (m *Microservice) Listen(endpoint string, callback func(message *Message)) {
	m.client.On(endpoint, callback)
}

// Call sends a request to the microservice
func (m *Microservice) Call(endpoint string, message *Message) (*Message, error) {
	msResponse := new(Message)

	if duration, err := time.ParseDuration("2s"); err != nil {
		return nil, err
	} else if raw, err := m.client.Ack(endpoint, message, duration); err != nil {
		return nil, err
	} else if err := json.Unmarshal([]byte(raw), msResponse); err != nil {
		return nil, err
	}

	return msResponse, nil
}

// RegisterMicroservice configures, automatically connects, and adds the microservice to App.Microservices
func RegisterMicroservice(name string, host string, port int, secure bool) error {
	App.Microservices[name] = new(Microservice)
	App.Microservices[name].configure(host, port, secure)
	_, err := App.Microservices[name].Connect()
	return err
}
