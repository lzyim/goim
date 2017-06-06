package message

import (
	"encoding/json"
	"github.com/satori/go.uuid"
)

type Message struct {
	Fuuid uuid.UUID
	Tuuid uuid.UUID
	Data  []byte
}

func (msg *Message) Unmarshal(data []byte, v *Message) error {
	return json.Unmarshal(data, v)
}

func (msg *Message) Marshal() ([]byte, error) {
	return json.Marshal(msg)
}
