package gosf

// OnConnect registers an event handler that fires when a client connects
func OnConnect(callback func(client *Client, request *Request)) {
	on("connect", func(args ...interface{}) {
		if len(args) > 1 {
			callback(args[0].(*Client), args[1].(*Request))
		}
	})
}

// OnDisconnect registers an event handler that fires when a client disconnects
func OnDisconnect(callback func(client *Client, request *Request)) {
	on("disconnect", func(args ...interface{}) {
		if len(args) > 1 {
			callback(args[0].(*Client), args[1].(*Request))
		}
	})
}

// OnBeforeRequest registers an event handler that fires before a request is processed by the controller
func OnBeforeRequest(callback func(client *Client, request *Request)) {
	on("before-request", func(args ...interface{}) {
		if len(args) > 1 {
			callback(args[0].(*Client), args[1].(*Request))
		}
	})
}

// OnAfterRequest registers an event handler that fires after a request is processed by the controller
func OnAfterRequest(callback func(client *Client, request *Request, response *Message)) {
	on("after-request", func(args ...interface{}) {
		if len(args) > 2 {
			callback(args[0].(*Client), args[1].(*Request), args[2].(*Message))
		}
	})
}

// OnBeforeResponse registers an event handler that fires before a response is processed by the controller
func OnBeforeResponse(callback func(client *Client, request *Request, response *Message)) {
	on("before-response", func(args ...interface{}) {
		if len(args) > 2 {
			callback(args[0].(*Client), args[1].(*Request), args[2].(*Message))
		}
	})
}

// OnAfterResponse registers an event handler that fires after a response is processed by the controller
func OnAfterResponse(callback func(client *Client, request *Request, response *Message)) {
	on("after-response", func(args ...interface{}) {
		if len(args) > 2 {
			callback(args[0].(*Client), args[1].(*Request), args[2].(*Message))
		}
	})
}

// OnBeforeClientBroadcast registers an event handler that fires before a client broadcast is sent to other clients
func OnBeforeClientBroadcast(callback func(client *Client, endpoint string, room string, response *Message)) {
	on("before-client-broadcast", func(args ...interface{}) {
		if len(args) > 3 {
			callback(args[0].(*Client), args[1].(string), args[2].(string), args[3].(*Message))
		}
	})
}

// OnAfterClientBroadcast registers an event handler that fires after a client broadcast is sent to other clients
func OnAfterClientBroadcast(callback func(client *Client, endpoint string, room string, response *Message)) {
	on("after-client-broadcast", func(args ...interface{}) {
		if len(args) > 3 {
			callback(args[0].(*Client), args[1].(string), args[2].(string), args[3].(*Message))
		}
	})
}

// OnBeforeBroadcast registers an event handler that fires before a global broadcast is sent to all clients
func OnBeforeBroadcast(callback func(endpoint string, room string, response *Message)) {
	on("before-broadcast", func(args ...interface{}) {
		if len(args) > 2 {
			callback(args[0].(string), args[1].(string), args[2].(*Message))
		}
	})
}

// OnAfterBroadcast registers an event handler that fires after a global broadcast is sent to all clients
func OnAfterBroadcast(callback func(endpoint string, room string, response *Message)) {
	on("after-broadcast", func(args ...interface{}) {
		if len(args) > 2 {
			callback(args[0].(string), args[1].(string), args[2].(*Message))
		}
	})
}

var hooks map[string][]func(...interface{})

func init() {
	hooks = make(map[string][]func(...interface{}))
}

func on(event string, callback func(...interface{})) {
	if hooks[event] == nil {
		hooks[event] = make([]func(...interface{}), 0)
	}

	hooks[event] = append(hooks[event], callback)
}

func emit(event string, args ...interface{}) {
	if hooks[event] == nil {
		return
	}

	for i := range hooks[event] {
		hooks[event][i](args...)
	}
}
