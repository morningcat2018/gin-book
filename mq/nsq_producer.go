package mq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func init() {
	nsqAddress := "127.0.0.1:4150"
	config := nsq.NewConfig()
	producer, _ = nsq.NewProducer(nsqAddress, config)
	return
}

func SendMqMessage(topicName string, byteData []byte) {
	// data string -> []byte(data)
	err := producer.Publish(topicName, byteData)
	if err != nil {
		fmt.Printf("publish msg to nsq failed, err:%v\n", err)
	}
}
