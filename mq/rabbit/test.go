package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type Publisher interface {
	Publish(ctx context.Context, msg interface{}) error
}

type Consumer func(ctx context.Context, message []byte) error

type Subscriber interface {
	Subscribe(consumer Consumer) error
}

type RabbitMqConfiguration struct {
	addr     string
	port     int
	user     string
	password string
	mutex    sync.Mutex
	conn     *amqp.Connection
}

func (c *RabbitMqConfiguration) Url() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.user, c.password, c.addr, c.port)
}

// NewRabbitMqConfiguration 创建RabbitMQ配置对象
func NewRabbitMqConfiguration(addr string, port int, user, password string) (*RabbitMqConfiguration, func(), error) {
	config := &RabbitMqConfiguration{
		addr:     addr,
		port:     port,
		user:     user,
		password: password,
	}
	conn, err := amqp.Dial(config.Url())
	if err != nil {
		return nil, func() {}, err
	}
	config.mutex = sync.Mutex{}
	config.conn = conn
	return config, func() {
		_ = conn.Close()
	}, nil
}

var _ Publisher = (*rabbitMqChannel)(nil)
var _ Subscriber = (*rabbitMqChannel)(nil)

type rabbitMqChannel struct {
	config            *RabbitMqConfiguration
	mutex             sync.Mutex
	closed            bool
	channel           *amqp.Channel
	exchange          string
	key               string
	queue             string
	channelClosedChan chan *amqp.Error
	done              chan struct{}
	deliveries        <-chan amqp.Delivery
	consumer          Consumer
	reconsume         chan struct{}
}

func NewRabbitMQPublisher(config *RabbitMqConfiguration, exchange, key, queue string) (Publisher, func(), error) {
	config.mutex.Lock()
	ch, err := config.conn.Channel()
	if err != nil {
		return nil, func() {}, err
	}
	config.mutex.Unlock()
	channel := &rabbitMqChannel{
		config:            config,
		mutex:             sync.Mutex{},
		channel:           ch,
		exchange:          exchange,
		key:               key,
		queue:             queue,
		channelClosedChan: make(chan *amqp.Error),
		done:              make(chan struct{}),
	}
	clean := func() {
		channel.mutex.Lock()
		defer channel.mutex.Unlock()
		channel.closed = true
		_ = channel.channel.Close()
		close(channel.done)
	}
	err = declareDirectExchange(ch, exchange)
	if err != nil {
		return nil, clean, err
	}
	err = declareAndBindQueue(ch, exchange, key, queue)
	if err != nil {
		return nil, clean, err
	}
	if err = ch.Confirm(false); err != nil {
		return nil, clean, err
	}
	//channel.confirmCh = ch.NotifyPublish(make(chan amqp.Confirmation))
	//ch.NotifyReturn()
	ch.NotifyClose(channel.channelClosedChan)
	go channel.keepalive()
	return channel, clean, nil
}

func NewRabbitMqSubscriber(config *RabbitMqConfiguration, exchange, key, queue string) (Subscriber, func(), error) {
	config.mutex.Lock()
	ch, err := config.conn.Channel()
	if err != nil {
		return nil, func() {}, err
	}
	config.mutex.Unlock()
	clean := func() {
		_ = ch.Close()
	}
	channel := &rabbitMqChannel{
		config:            config,
		mutex:             sync.Mutex{},
		channel:           ch,
		exchange:          exchange,
		key:               key,
		queue:             queue,
		channelClosedChan: make(chan *amqp.Error),
		done:              make(chan struct{}),
	}
	clean = func() {
		channel.mutex.Lock()
		defer channel.mutex.Unlock()
		channel.closed = true
		_ = channel.channel.Close()
		close(channel.done)
	}
	err = declareDirectExchange(ch, exchange)
	if err != nil {
		return nil, clean, err
	}
	err = declareAndBindQueue(ch, exchange, key, queue)
	if err != nil {
		return nil, clean, err
	}
	if err = ch.Confirm(false); err != nil {
		return nil, clean, err
	}
	//channel.confirmCh = ch.NotifyPublish(make(chan amqp.Confirmation))
	//ch.NotifyReturn()
	ch.NotifyClose(channel.channelClosedChan)
	go channel.keepalive()
	return channel, clean, nil
}

func (ch *rabbitMqChannel) Publish(ctx context.Context, msg interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err = ch.ensureChannelOpened(); err != nil {
		return err
	}

	err = ch.channel.PublishWithContext(
		ctx,
		ch.exchange,
		ch.key,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		})
	if err != nil {
		return err
	}

	return nil
}

func (ch *rabbitMqChannel) Subscribe(consumer Consumer) error {
	deliveries, err := ch.channel.ConsumeWithContext(context.Background(), ch.queue, "", false, false, false, false, nil)
	if err != nil {
		return err
	}
	ch.deliveries = deliveries
	ch.consumer = consumer
	ch.reconsume = make(chan struct{})
	go func() {
		for {
			select {
			case <-ch.done:
				close(ch.reconsume)
				fmt.Printf("subscriber: %s, exited\n", ch.queue)
				return
			case <-ch.reconsume:
				ch.deliveries, _ = ch.channel.ConsumeWithContext(context.Background(), ch.queue, "", false, false, false, false, nil)
			case delivery := <-ch.deliveries:
				ch.consume(delivery)
			}
		}
	}()
	return nil
}

func (ch *rabbitMqChannel) consume(delivery amqp.Delivery) {
	defer func() {
		if err := recover(); err != nil {
			_ = delivery.Nack(false, true)
		}
	}()
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := ch.consumer(timeout, delivery.Body)
	if err != nil {
		fmt.Printf("subscriber: %s, deliveryTag: %d, msg: %s, err: %v\n", ch.queue, delivery.DeliveryTag, delivery.Body, err)
		_ = delivery.Nack(false, true)
		return
	}
	_ = delivery.Ack(false)
}

func (ch *rabbitMqChannel) keepalive() {
	for {
		select {
		case <-ch.done:
			fmt.Println("RabbitMq channel exited")
			return
		case err := <-ch.channelClosedChan:
			if err == nil && !ch.channel.IsClosed() {
				continue
			}
			ch.mutex.Lock()
			if ch.closed == true {
				ch.mutex.Unlock()
				continue
			}
			ch.mutex.Unlock()
			fmt.Printf("channel closed: %v\n", err)
			connectErr := ch.ensureChannelOpened()
			if connectErr != nil {
				fmt.Printf("reconnect error: %v", connectErr)
				time.Sleep(5 * time.Second)
				continue
			}
			if ch.consumer != nil {
				ch.reconsume <- struct{}{}
			}
		}
	}

}

func (ch *rabbitMqChannel) ensureChannelOpened() (err error) {
	if !ch.channel.IsClosed() {
		return
	}
	ch.mutex.Lock()
	defer ch.mutex.Unlock()
	if !ch.channel.IsClosed() {
		return nil
	}
	config := ch.config
	config.mutex.Lock()
	defer config.mutex.Unlock()
	if config.conn.IsClosed() {
		if config.conn, err = amqp.Dial(config.Url()); err != nil {
			return err
		}
		fmt.Printf("reconnected to RabbitMQ, url: %s\n", config.Url())
	}
	if ch.channel, err = ch.config.conn.Channel(); err != nil {
		return err
	}
	fmt.Printf("reopen channel for queue: %s\n", ch.queue)
	return nil
}

func declareDirectExchange(ch *amqp.Channel, exchange string) error {
	if err := ch.ExchangeDeclare(exchange, amqp.ExchangeDirect, true, false, false, false, nil); err != nil {
		return err
	}
	fmt.Printf("declared exchange, name: %s\n", exchange)
	return nil
}

func declareAndBindQueue(ch *amqp.Channel, exchange, key, queue string) error {
	q, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	fmt.Printf("declared queue, name: %s, messages: %d, consumers: %d\n", q.Name, q.Messages, q.Consumers)
	err = ch.QueueBind(queue, key, exchange, false, nil)
	if err == nil {
		fmt.Printf("bind queue to exchange, exchange: %s, key: %s, queue: %s\n", exchange, key, queue)
	}
	return err
}

func TestPublish() {
	conn, clean, err := NewRabbitMqConfiguration("192.168.204.134", 5672, "emove", "668466")
	if err != nil {
		clean()
		panic(err)
	}
	publisher, f, err := NewRabbitMQPublisher(conn, "test", "", "test")
	if err != nil {
		clean()
		f()
		panic(err)
	}
	start := time.Now()
	for i := 1; i <= 5; i++ {
		err = publisher.Publish(context.Background(), map[string]interface{}{"msgId": i})
		if err != nil {
			clean()
			f()
			panic(err)
		}
		time.Sleep(time.Second * 1)
	}
	println("elapsed: \n", time.Since(start).Seconds())
	f()
}

func TestSubscribe() {
	conn, clean, err := NewRabbitMqConfiguration("192.168.204.134", 5672, "emove", "668466")
	if err != nil {
		clean()
		panic(err)
	}
	subscriber, f, err := NewRabbitMqSubscriber(conn, "test", "", "test")
	if err != nil {
		clean()
		f()
		panic(err)
	}
	consumeTime := time.Now()
	err = subscriber.Subscribe(func(ctx context.Context, message []byte) error {
		fmt.Printf("consume: %s\n", string(message))
		consumeTime = time.Now()
		return nil
	})
	if err != nil {
		panic(err)
	}
	for time.Since(consumeTime) < 5*time.Second {
		time.Sleep(time.Second)
	}
	f()
}
