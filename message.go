package gosf

// Message - Standard message type for Socket communications
type Message struct {
	ID   int    `json:"id,omitempty"`
	GUID string `json:"guid,omitempty"`
	UUID string `json:"uuid,omitempty"`

	Success bool   `json:"success"`
	Text    string `json:"text,omitempty"`

	Meta map[string]interface{} `json:"meta,omitempty"`
	Body map[string]interface{} `json:"body,omitempty"`
}

// WithoutMeta removes meta information before returning a copy of the message
func (m *Message) WithoutMeta() *Message {
	if m.Meta == nil {
		return m
	}

	var metaFreeMessage = new(Message)
	(*metaFreeMessage) = *m
	metaFreeMessage.Meta = nil
	return metaFreeMessage
}

// NewSuccessMessage generates a success message
func NewSuccessMessage(args ...interface{}) *Message {
	m := new(Message)

	m.Success = true

	var text = ""
	var body = map[string]interface{}(nil)

	for i, p := range args {
		switch i {
		case 0: // text
			param, ok := p.(string)
			if !ok {
				break
			}
			text = param
		case 1: // body
			param, ok := p.(map[string]interface{})
			if !ok {
				break
			}
			body = param
		}
	}

	if text != "" {
		m.Text = text
	}
	m.Body = body

	return m
}

// NewFailureMessage generates a failure message
func NewFailureMessage(args ...interface{}) *Message {
	m := new(Message)

	m.Success = false

	var text = ""
	var body = map[string]interface{}(nil)

	for i, p := range args {
		switch i {
		case 0: // text
			param, ok := p.(string)
			if !ok {
				break
			}
			text = param
		case 1: // body
			param, ok := p.(map[string]interface{})
			if !ok {
				break
			}
			body = param
		}
	}

	if text != "" {
		m.Text = text
	}
	m.Body = body

	return m
}
