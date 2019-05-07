package gosf

import (
	"log"
	"testing"
	"time"
)

// MyPlugin is the aspect oriented element required by the modular plugin framework
type MyPlugin struct{}

// Activate is an aspect-oriented modular plugin requirement
func (p MyPlugin) Activate(app *AppSettings) {}

// Deactivate is an aspect-oriented modular plugin requirement
func (p MyPlugin) Deactivate(app *AppSettings) {}

func TestRegisterPlugin(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Unable to register plugin.")
		}
	}()
	RegisterPlugin(new(MyPlugin))
}

func endpointEcho(client *Client, request *Request) *Message {
	return NewSuccessMessage(request.Message.Text, nil)
}

func TestRegisterEndpoint(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Echo endpoint could not be registered.")
		}
	}()

	Listen("echo", endpointEcho)
}

func TestStartServer(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Unable to start server.")
		}
	}()

	// Start server in another thread
	go Startup(map[string]interface{}{
		"port": 9988,
	})
	time.Sleep(2 * time.Second)
}

func TestRegisterMicroservice(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Unable to register microservice.")
		}
	}()

	// Register a microservice connection to the server started above
	RegisterMicroservice("utils", "127.0.0.1", 9988, false)

	// Wait for server microservice to connect
	time.Sleep(2 * time.Second)
}

func TestCallMicroserviceMessage(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Echo endpoint could not be called successfully.")
		}
	}()

	input := NewSuccessMessage("Hello world.", nil)
	response, err := App.Microservices["utils"].Call("echo", input)

	if err != nil {
		panic(err)
	}

	if response.Text == "" || response.Text != "Hello world." {
		panic("Response text did not return as expected.")
	}
}

func TestLobMicroserviceMessage(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Echo endpoint could not be called successfully.")
		}
	}()

	input := NewSuccessMessage("Hello world.", nil)
	err := App.Microservices["utils"].Lob("echo", input)

	if err != nil {
		panic(err)
	}
}
