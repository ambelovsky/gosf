package gosf

import (
	"reflect"
	"strings"
	"unicode"
)

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
func NewSuccessMessage(text string, body interface{}) *Message {
	m := new(Message)

	m.Success = true
	if text != "" {
		m.Text = text
	}
	m.Body = convertToJSONMap(body)

	return m
}

func convertToJSONMap(data interface{}) map[string]interface{} {
	jsonMap := make(map[string]interface{})

	arg := reflect.ValueOf(data).Elem()
	for j := 0; j < arg.NumField(); j++ {
		value := arg.Field(j).Interface()
		valueType := arg.Field(j).Kind()

		if value == nil {
			continue
		}

		name := arg.Type().Field(j).Name
		nameBytes := []byte(name)
		nameBytes[0] = byte(unicode.ToLower(rune(nameBytes[0])))
		name = string(nameBytes)

		tag := arg.Type().Field(j).Tag.Get("json")
		if tag != "" {
			tagParts := strings.Split(tag, ",")
			if len(tagParts) > 0 {
				name = tagParts[0]
			}

			// ignore requests for omitempty, this function omits empty by default
			// no check for other tag parts
		}

		if valueType == reflect.Struct {
			jsonMap[name] = convertToJSONMap(&value)
		} else {
			jsonMap[name] = value
		}
	}
	return jsonMap
}
