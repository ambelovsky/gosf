package gosf

// Message - Standard message type for Socket communications
type Message struct {
	ID int `json:"id,omitempty"`

	Success bool   `json:"success"`
	Text    string `json:"text,omitempty"`

	Meta interface{} `json:"meta,omitempty"`
	Body interface{} `json:"body,omitempty"`
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
func NewSuccessMessage(text string, body interface{}) *Message {
	m := new(Message)

	m.Success = true
	if text != "" {
		m.Text = text
	}
	m.Body = body

	return m
}
