package gosf

import (
	"log"
	"strconv"
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
			log.Println(err)
			t.Error("echo endpoint could not be called successfully")
		}
	}()

	input := NewSuccessMessage("Hello world.")

	ms := GetMicroservice("utils")
	if ms == nil {
		panic("no microservice was returned by GetMicroservice")
	}

	var response *Message
	var err error

	start := time.Now()
	for i := 0; i < 2; i++ {
		response, err = ms.Call("echo", input)
	}
	elapsed := time.Since(start)

	log.Printf(" - 2 calls in %s", elapsed)

	if err != nil {
		panic(err)
	}

	if response.Text == "" || response.Text != "Hello world." {
		panic("response text did not return as expected")
	}
}

func TestGoMicroserviceMessage(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
			t.Error("echo endpoint could not be called successfully")
		}
	}()

	ms := GetMicroservice("utils")
	if ms == nil {
		panic("no microservice was returned by GetMicroservice")
	}

	var response *Message
	var err error
	var chMsg []chan *GoMessage
	var numberOfCalls = 1000

	start := time.Now()
	for i := 0; i < numberOfCalls; i++ {
		input := NewSuccessMessage("Hello world " + strconv.Itoa(i+1))
		chMsg = append(chMsg, ms.Go("echo", input))
	}

	for i := 0; i < len(chMsg); i++ {
		response, err = ReadGoMessage(chMsg[i])
		close(chMsg[i])
	}
	elapsed := time.Since(start)

	log.Printf(" - "+strconv.Itoa(numberOfCalls)+" calls in %s", elapsed)

	if err != nil {
		panic(err)
	}

	if response.Text == "" {
		panic("response text did not return as expected")
	}
}

func TestLobMicroserviceMessage(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error(err)
			t.Error("Echo endpoint could not be called successfully.")
		}
	}()

	input := NewSuccessMessage("Hello world.", nil)

	var err error

	start := time.Now()
	for i := 0; i < 10; i++ {
		err = App.Microservices["utils"].Lob("echo", input)
	}
	elapsed := time.Since(start)

	log.Printf(" - 10 calls in %s", elapsed)

	if err != nil {
		panic(err)
	}
}
