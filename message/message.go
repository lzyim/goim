package message

import (
	"encoding/json"
)

type Message struct {
	Fuuid string
	Tuuid string
	Data  string
}

func (msg *Message) Unmarshal(data []byte, v *Message) error {
	return json.Unmarshal(data, v)
}

func (msg *Message) Marshal() ([]byte, error) {
	return json.Marshal(msg)
}
