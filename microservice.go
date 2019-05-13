package gosf

import (
	"encoding/json"
	"sync"
	"time"

	io "github.com/ambelovsky/gosf-socketio"
	transport "github.com/ambelovsky/gosf-socketio/transport"
)

var maxConnections int

type GoMessage struct {
	message *Message
	error   error
}

func init() {
	App.Microservices = make(map[string]*Microservice)
	maxConnections = 2
}

// Microservice defines a microservice connection
type Microservice struct {
	host            string
	port            int
	secure          bool
	transport       *transport.WebsocketTransport
	clients         []*io.Client
	nextClient      int
	nextClientMutex sync.Mutex
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
	if len(m.clients) < 1 {
		return false
	}

	connected := true
	for _, c := range m.clients {
		if !c.IsAlive() {
			connected = false
			break
		}
	}

	return connected
}

// Connect manually connects or reconnects to the microservice
func (m *Microservice) Connect() (*Microservice, error) {
	url := io.GetUrl(m.host, m.port, m.secure)

	if m.Connected() {
		m.Disconnect()
	}

	if len(m.clients) < 1 {
		for i := 0; i < maxConnections; i++ {
			newCli, err := io.Dial(url, m.transport)
			if err != nil {
				return m, err
			}
			m.clients = append(m.clients, newCli)
		}
	} else {
		for i := range m.clients {
			var err error
			m.clients[i], err = io.Dial(url, m.transport)
			if err != nil {
				return m, err
			}
		}
	}

	return m, nil
}

// Disconnect manually terminates the connection to the microservice
func (m *Microservice) Disconnect() {
	if m.Connected() == false {
		return
	}

	for i := range m.clients {
		m.clients[i].Close()
	}
}

// Listen registers and event handler for a microservice endpoint
func (m *Microservice) Listen(endpoint string, callback func(message *Message)) {
	for i := range m.clients {
		m.clients[i].On(endpoint, callback)
	}
}

func (m *Microservice) getNextClient() *io.Client {
	m.nextClientMutex.Lock()
	defer m.nextClientMutex.Unlock()

	m.nextClient++
	if m.nextClient >= maxConnections {
		m.nextClient = 0
	}

	return m.clients[m.nextClient]
}

// Call sends a request to the microservice
func (m *Microservice) Call(endpoint string, message *Message) (*Message, error) {
	msResponse := new(Message)

	if duration, err := time.ParseDuration("2s"); err != nil {
		return nil, err
	} else if raw, err := m.getNextClient().Ack(endpoint, message, duration); err != nil {
		return nil, err
	} else if err := json.Unmarshal([]byte(raw), msResponse); err != nil {
		return nil, err
	}

	err := recover()
	if err != nil {
		return nil, err.(error)
	}
	return msResponse, nil
}

// Go sends a request to the microservice returning channels
func (m *Microservice) Go(endpoint string, message *Message) chan *GoMessage {
	goChan := make(chan *GoMessage)

	go func(goChan chan *GoMessage) {
		msg, err := m.Call(endpoint, message)
		goMsg := &GoMessage{
			message: msg,
			error:   err,
		}
		goChan <- goMsg
	}(goChan)

	return goChan
}

// Lob sends a request to the microservice that does not require a response
func (m *Microservice) Lob(endpoint string, message *Message) error {
	err := m.getNextClient().Emit(endpoint, message)
	return err
}

// ReadGoMessage reads from a microservice Go channel of type *GoMessage
func ReadGoMessage(chMsg chan *GoMessage) (*Message, error) {
	msg := <-chMsg
	return msg.message, msg.error
}

// RegisterMicroservice configures, automatically connects, and adds the microservice to App.Microservices
func RegisterMicroservice(name string, host string, port int, secure bool) error {
	App.Microservices[name] = new(Microservice)
	App.Microservices[name].configure(host, port, secure)
	_, err := App.Microservices[name].Connect()
	return err
}

// DeregisterMicroservice disconnects and removes a microservice from App.Microservices
func DeregisterMicroservice(name string) {
	App.Microservices[name].Disconnect()
	delete(App.Microservices, name)
}

// GetMicroservice retrieves a microservice reference from App.Microservices
func GetMicroservice(name string) *Microservice {
	if microservice, ok := App.Microservices[name]; ok {
		return microservice
	}
	return nil
}
