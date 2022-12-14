package nsqer

import (
	"errors"
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

type nsqProducer struct {
	producer *nsq.Producer
}

func (p *nsqProducer) Publish(topic string, message []byte) (err error) {
	if len(message) == 0 {
		return errors.New("message is empty")
	}
	if err = p.producer.Publish(topic, message); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 延迟消息
func (p *nsqProducer) DeferredPublish(topic string, delay time.Duration, message []byte) (err error) {
	if len(message) == 0 {
		return errors.New("message is empty")
	}
	if err = p.producer.DeferredPublish(topic, delay, message); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func initProducer(addr string) (p *nsq.Producer, err error) {
	config := nsq.NewConfig()
	if p, err = nsq.NewProducer(addr, config); err != nil {
		return nil, err
	}
	return p, nil
}

var NsqProducer *nsqProducer

func init() {
	NsqProducer = &nsqProducer{}
	NsqProducer.producer, _ = initProducer("192.168.168.128:4150")
}
