package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	Nats  *nats.Conn
	Topic string
}

func NewMsgHandler(natsURL string) *MsgHandler {
	// Connect to a server
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to NATS Server: %s\n", natsURL)

	return &MsgHandler{
		Nats:  nc,
		Topic: "Test",
	}
}

func (m *MsgHandler) Handle(msg []byte) {
	log.Printf("publishing message to mq: %s", string(msg))
	m.Nats.Publish(m.Topic, msg)
}

func ConnectMQ() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	return nc
}
