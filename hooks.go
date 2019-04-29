package gosf

// OnConnect registers an event handler that fires when a client connects
func OnConnect(callback func(client *Client, request *Request)) {
	on("connect", func(args ...interface{}) {
		callback(args[0].(*Client), args[1].(*Request))
	})
}

// OnDisconnect registers an event handler that fires when a client disconnects
func OnDisconnect(callback func(client *Client, request *Request)) {
	on("disconnect", func(args ...interface{}) {
		callback(args[0].(*Client), args[1].(*Request))
	})
}

// OnBeforeRequest registers an event handler that fires before a request is processed by the controller
func OnBeforeRequest(callback func(client *Client, request *Request)) {
	on("before-request", func(args ...interface{}) {
		callback(args[0].(*Client), args[1].(*Request))
	})
}

// OnAfterRequest registers an event handler that fires after a request is processed by the controller
func OnAfterRequest(callback func(client *Client, request *Request, response *Message)) {
	on("after-request", func(args ...interface{}) {
		callback(args[0].(*Client), args[1].(*Request), args[2].(*Message))
	})
}

// OnBeforeResponse registers an event handler that fires before a response is processed by the controller
func OnBeforeResponse(callback func(client *Client, request *Request, response *Message)) {
	on("before-response", func(args ...interface{}) {
		callback(args[0].(*Client), args[1].(*Request), args[2].(*Message))
	})
}

// OnAfterResponse registers an event handler that fires after a response is processed by the controller
func OnAfterResponse(callback func(client *Client, request *Request, response *Message)) {
	on("after-response", func(args ...interface{}) {
		callback(args[0].(*Client), args[1].(*Request), args[2].(*Message))
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
