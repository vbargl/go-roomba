package roomba

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type stateMessage struct {
	State reportedMessage `json:"state"`
}

type reportedMessage struct {
	Reported map[string]any `json:"reported"`
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
	}

	r.statusMutex.Lock()
	statusValue := reflect.ValueOf(r.status).Elem()
	for field, value := range reportedState.State.Reported {
		f := getFieldByTag(field, statusValue)
		if f.IsValid() {
			f.Set(reflect.ValueOf(value))
		} else if r.debug {
			log.Printf("Field %s not found in status struct", field)
		}
	}
	r.statusMutex.Unlock()
}

func getFieldByTag(tag string, v reflect.Value) reflect.Value {
	for i := 0; i < v.NumField(); i++ {
		tagContent := v.Type().Field(i).Tag.Get("json")
		if strings.Split(tagContent, ",")[0] == tag {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}
