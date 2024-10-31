package roomba

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type stateMessage struct {
	State reportedMessage `json:"state"`
}

type reportedMessage struct {
	Reported json.RawMessage `json:"reported"`
}

func (r *Roomba) stateMessageHandler(client mqtt.Client, msg mqtt.Message) {
	if r.debug {
		log.Printf("RECV: %s", string(msg.Payload()))
	}

	if r.stateWriter != nil {
		r.stateWriter.Write(append(msg.Payload(), byte('\n')))
	}

	reportedState := stateMessage{}
	err := json.Unmarshal(msg.Payload(), &reportedState)
	if err != nil {
		log.Printf("Roomba state unmarshal error: %s", err.Error())
		return
	}

	r.statusMutex.Lock()
	defer r.statusMutex.Unlock()
	err = json.Unmarshal(reportedState.State.Reported, r.status)
	if err != nil {
		log.Printf("Roomba status unmarshal error: %s", err.Error())
	}
}
