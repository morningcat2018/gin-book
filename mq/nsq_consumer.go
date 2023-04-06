package mq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

// 消费者
const topic = "topic_book"
const channel = "topic_book_0"
const nsqd_address = "127.0.0.1:4150"
const lookupd_address = "127.0.0.1:4161"

func init() {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer := &BookHandler{
		Title: "book-service",
	}
	c.AddHandler(consumer)

	// if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
	if err := c.ConnectToNSQLookupd(lookupd_address); err != nil { // 通过lookupd查询
		panic(err.Error())
	}
	return
}

// Handler 是一个消费者类型
type BookHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *BookHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}
