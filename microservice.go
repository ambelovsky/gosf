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

func (m *Microservice) Connected() bool {
	if m.client == nil {
		return false
	}
	return m.client.IsAlive()
}

func (m *Microservice) Connect() (*Microservice, error) {
	if m.Connected() {
		m.Disconnect()
	}

	var err error
	url := io.GetUrl(m.host, m.port, m.secure)
	m.client, err = io.Dial(url, m.transport)
	return m, err
}

func (m *Microservice) Disconnect() {
	if m.Connected() == false {
		return
	}
	m.client.Close()
}

func (m *Microservice) Call(endpoint string, message Message) (*Message, error) {
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

func RegisterMicroservice(name string, host string, port int, secure bool) {
	App.Microservices[name] = new(Microservice)
	App.Microservices[name].configure(host, port, secure)
	go App.Microservices[name].Connect()
}
