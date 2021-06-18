package listener

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	nats "github.com/nats-io/nats.go"
)

//go:generate  easyjson nats_publisher.go

// NatsPublisher represent event publisher.
type NatsPublisher struct {
	conn *nats.Conn
}

// Close NATS connection.
func (n NatsPublisher) Close() error {
	n.conn.Close()
	return nil
}

// Event event structure for publishing to the NATS server.
//easyjson:json
type Event struct {
	ID        uuid.UUID              `json:"id"`
	Schema    string                 `json:"schema"`
	Table     string                 `json:"table"`
	Action    string                 `json:"action"`
	Data      map[string]interface{} `json:"data"`
	EventTime time.Time              `json:"commitTime"`
}

// Publish serializes the event and publishes it on the bus.
func (n NatsPublisher) Publish(subject string, event Event) error {
	msg, err := event.MarshalJSON()
	if err != nil {
		return fmt.Errorf("marshal err: %w", err)
	}
	return n.conn.Publish(subject, msg)
}

// NewNatsPublisher return new NatsPublisher instance.
func NewNatsPublisher(conn *nats.Conn) *NatsPublisher {
	return &NatsPublisher{conn: conn}
}

// GetSubjectName creates subject name from the prefix, schema and table name.
func (e Event) GetSubjectName(prefix string) string {
	if prefix == "" {
		return fmt.Sprintf("wal.%s.%s", e.Schema, e.Table)
	} else {
		return fmt.Sprintf("%s.%s.%s", prefix, e.Table, strings.ToLower(e.Action))
	}
}
