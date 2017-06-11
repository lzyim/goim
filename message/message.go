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

func (msg *Message) Unmarshal(data []byte) error {
	return json.Unmarshal(data, msg)
}

func (msg *Message) Marshal() ([]byte, error) {
	return json.Marshal(msg)
}
