package mq

import (
	"context"

	"github.com/tork/task"
)

// Broker is the message-queue, pub/sub mechanism used for delivering tasks.
type Broker interface {
	Queues(ctx context.Context) []string
	Publish(ctx context.Context, qname string, t *task.Task) error
	Subscribe(qname string, handler func(ctx context.Context, t *task.Task) error) error
}