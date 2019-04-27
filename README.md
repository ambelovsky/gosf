# gosf
GoLang SocketIO Framework or GOSF is an easy-to-use framework for developing SocketIO API's in GoLang. An example server to help you get started can be found at [github.com/ambelovsky/gosf-sample-app](https://github.com/ambelovsky/gosf-sample-app).

For an in-depth look at the API Framework, check us out at [gosf.io](http://gosf.io).

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

func echo(request *f.Request) {
  response := new(f.Message)
  response.Success = true
  response.Text = request.Message.Text

  request.Respond(response)
}

func init() {
  // Listen on an endpoint
  f.Listen("echo", echo)
}

func main() {
  // Start the server using a basic configuration
  f.Startup(map[string]interface{}{
    "port": 9999})
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

## Learn More

Discover more about GOSF at [gosf.io](http://gosf.io).

## Documenting Your API

While you're building your API, take some time to build the documentation too!  Check out [github.com/ambelovsky/go-api-docs](https://github.com/ambelovsky/go-api-docs) for an
easy-to-use documentation system built using the slate theme.

## Original Author

- [Aaron Belovsky](https://github.com/ambelovsky)

## License

MIT
