# gosf
GoLang SocketIO Framework or GOSF is an easy-to-use framework for developing SocketIO API's in GoLang. An example server to help you get started can be found
at [github.com/ambelovsky/gosf-sample-app](https://github.com/ambelovsky/gosf-sample-app).

## Get It

```sh
go get -u "github.com/ambelovsky/gosf"
```

## Quick Start

The following sample will start a server that responds on an "echo" endpoint and return the same message received from the client back to the client.

### Server

```go
package main

import (
	f "github.com/ambelovsky/gosf"
)

func echo(request *f.Request, clientMessage *f.Message) {
	response := new(f.Message)
	response.Success = true
	response.Message = clientMessage.Message

	request.Respond(clientMessage, response)
}

func init() {
	// Load server config file
  f.LoadConfig("server", "server.json")
  
  // Listen on an endpoint
  f.Listen("echo", echo)
}

func main() {
	// Start the server using the loaded configuration
	f.Startup(f.Config["server"].(map[string]interface{}))
}
```

### Client
```html
<script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/2.0.3/socket.io.slim.js"></script>
<script>
  var socket = io.connect('ws://localhost:9999', { transports: ['websocket'] });

  socket.on('echo', function(response) {
    console.log(response);
  });

  socket.emit('echo', { message: 'Hello world.' });
</script>
```

## Original Author

- [Aaron Belovsky](https://github.com/ambelovsky)

## License

MIT
