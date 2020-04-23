package messages

import (
	"encoding/json"
	"fmt"
)

type MessageType string


type Message map[MessageType]interface{}

func (m *Message) UnmarshalJSON(b []byte) error {
	msgSub := make(map[MessageType]interface{})
	if err := json.Unmarshal(b, &msgSub); err != nil {
		return err
	}

	for key, val := range msgSub {
		data, err := json.Marshal(val)
		if err != nil {
			return err
		}
		msgSub[key] = data
	}
	*m = msgSub
	return nil
}

func MarshalMessage(msg interface{}) ([]byte, error) {
	retMsg := make(map[MessageType]interface{})

	switch msg.(type) {
	default:
		return nil, fmt.Errorf("Message type not recognized")
	}
	return json.Marshal(retMsg)
}

func UnmarshalMessage(data []byte) (MessageType, interface{}, error) {
	var msg Message
	var retMsg interface{}
	if err := json.Unmarshal(data, &msg); err != nil {
		return "", nil, err
	}

	for msgType, payload := range msg {

		switch msgType {
		default:
			return msgType, nil, nil
		}
		if err := json.Unmarshal(payload.([]byte), retMsg); err != nil {
			return "", nil, err
		}
		return msgType, retMsg, nil
	}
	return "", nil, fmt.Errorf("bad message")
}
