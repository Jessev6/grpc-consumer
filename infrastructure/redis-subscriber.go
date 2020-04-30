package infrastructure

import (
	"errors"
	"log"

	"github.com/go-redis/redis/v7"
	"github.com/verbruggenjesse/grpc-consumer/domain"
	"github.com/verbruggenjesse/grpc-consumer/domain/abstract"
)

// RedisSubscriber is used to subscribe to messages from redis
type RedisSubscriber struct {
	client *redis.Client
}

// NewRedisSubscriber is the constructer for RedisSubscriber
func NewRedisSubscriber(opts *redis.Options) (*RedisSubscriber, error) {
	client := redis.NewClient(opts)

	if err := testClient(client); err != nil {
		return nil, err
	}

	return &RedisSubscriber{
		client: client,
	}, nil
}

func testClient(c *redis.Client) error {
	_, err := c.Ping().Result()

	return err
}

func (r *RedisSubscriber) isInitialized() (bool, error) {
	if r.client == nil {
		return false, errors.New("Redis client was not initialized")
	}
	return true, nil
}

// Subscribe is used to subscribe to messages
func (r *RedisSubscriber) Subscribe(subscription abstract.ISubscription, messageChan *chan abstract.IMessage, errChan *chan error) {
	var eventSubscription *domain.EventSubscription
	ok := true

	if eventSubscription, ok = subscription.(*domain.EventSubscription); !ok {
		*errChan <- errors.New("could not validate requested subscription")
	}

	initialized, err := r.isInitialized()

	if !initialized {
		*errChan <- err
	}

	closedLength := false

	if subscription.Count() != 0 {
		closedLength = true
	}

	var messages []redis.XMessage

	if closedLength {
		if !eventSubscription.Reversed() {
			messages, err = r.client.XRangeN(subscription.Key(), subscription.From(), subscription.To(), int64(subscription.Count())).Result()

			if err != nil {
				*errChan <- err
			}
		} else {
			messages, err = r.client.XRevRangeN(subscription.Key(), subscription.From(), subscription.To(), int64(subscription.Count())).Result()

			if err != nil {
				*errChan <- err
			}
		}

		for _, message := range messages {
			log.Printf("event: %v", message)
			*messageChan <- domain.NewRedisMessage(subscription.Key(), message.Values)
		}

	} else {
		for {
			results, err := r.client.XRead(&redis.XReadArgs{
				Streams: []string{subscription.Key(),  "$"},
				Block:   0,
			}).Result()

			if err != nil {
				*errChan <- err
				break
			}

			res := results[0]

			msg := res.Messages[0]

			message := domain.NewRedisMessage(
				subscription.Key(),
				msg.Values,
			)

			*messageChan <- message
		}
	}
}
