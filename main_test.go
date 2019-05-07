package gosf

import (
	"log"
	"testing"
	"time"

	"github.com/ambelovsky/gosf"
)

func endpointEcho(client *gosf.Client, request *gosf.Request) *gosf.Message {
	return gosf.NewSuccessMessage(request.Message.Text, nil)
}

func TestRegisterEndpoint(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Echo endpoint could not be registered.")
		}
	}()

	gosf.Listen("echo", endpointEcho)
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
	go gosf.Startup(map[string]interface{}{
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
	gosf.RegisterMicroservice("utils", "127.0.0.1", 9988, false)

	// Wait for server microservice to connect
	time.Sleep(2 * time.Second)
}

func TestSendMicroserviceMessage(t *testing.T) {
	log.Println(t.Name())
	defer func() {
		err := recover()
		if err != nil {
			t.Error("Echo endpoint could not be called successfully.")
		}
	}()

	input := gosf.NewSuccessMessage("Hello world.", nil)
	response, err := gosf.App.Microservices["utils"].Call("echo", input)

	if err != nil {
		panic(err)
	}

	if response.Text == "" || response.Text != "Hello world." {
		panic("Response text did not return as expected.")
	}
}
