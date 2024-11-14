package natsutils

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

// Provider interface defines the basic operations that can be performed with NATS, such as publishing, subscribing, and cleaning up resources.
type Provider interface {
	Cleanup() error                                                                                // Cleanup gracefully shuts down the provider.
	Publish(subj string, data []byte) error                                                        // Publish sends a message to the specified subject.
	QueueSubscribe(subj string, queue string, handler nats.MsgHandler) (*nats.Subscription, error) // QueueSubscribe subscribes to a subject with a queue group.
}

// NatsProvider encapsulates a NATS connection, providing methods for publishing messages,
// subscribing to subjects, and managing connection cleanup.
type NatsProvider struct {
	natsConn *nats.Conn
}

// NewProvider initializes and returns a new Provider instance with the given NATS connection.
func NewProvider(natsConn *nats.Conn) *NatsProvider {
	return &NatsProvider{
		natsConn: natsConn,
	}
}

// Cleanup gracefully drains and closes the NATS connection, ensuring that all pending messages are sent and resources are released.
func (prov *NatsProvider) Cleanup() error {
	err := prov.natsConn.Drain()
	if err != nil {
		return fmt.Errorf("error closing nats connection: %w", err)
	}

	return nil
}

// Publish publishes the data argument to the given subject.
func (prov *NatsProvider) Publish(subj string, data []byte) error {
	return prov.natsConn.Publish(subj, data)
}

// Subscribe creates an asynchronous queue subscriber on the given subject and queue.
func (prov *NatsProvider) QueueSubscribe(subj string, queue string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return prov.natsConn.QueueSubscribe(subj, queue, handler)
}
