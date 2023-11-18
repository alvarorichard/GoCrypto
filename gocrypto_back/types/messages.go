package types

import "encoding/json"

type Message struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (m *Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
