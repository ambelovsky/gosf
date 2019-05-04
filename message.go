package gosf

// Message - Standard message type for Socket communications
type Message struct {
	ID int `json:"id,omitempty"`

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
func NewSuccessMessage(text string, body map[string]interface{}) *Message {
	m := new(Message)

	m.Success = true
	if text != "" {
		m.Text = text
	}
	m.Body = body

	return m
}

// NewFailureMessage generates a failure message
func NewFailureMessage(text string) *Message {
	m := new(Message)

	m.Success = false
	if text != "" {
		m.Text = text
	}

	return m
}
