package broker

import (
	"alior-manager-bot/src/config"
	"github.com/EgorKo25/common/broker"
	"github.com/EgorKo25/common/logger"
)

func InitBroker(brokerCfg config.BrokerConfig, log logger.ILogger) error {
	err := broker.InitBroker(brokerCfg.URL)
	if err != nil {
		log.Fatal(err.Error())
	}

	// паблишер для отправки ответов
	//pubConfig := broker.NewPublisherConfig(brokerCfg.Exchange.Name, brokerCfg.Exchange.Kind, "ans")
	//if pubConfig == nil {
	//	log.Fatal("publisher config not created")
	//}
	//
	//err = broker.CreatePublisher("ans_publisher", pubConfig)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	// паблишер для отправки запросов
	pubConfig := broker.NewPublisherConfig(brokerCfg.Exchange.Name, brokerCfg.Exchange.Kind, "ask")
	if pubConfig == nil {
		log.Fatal("publisher config not created")
	}

	err = broker.CreatePublisher("ask_publisher", pubConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// читатель ответов
	consConfig := broker.NewConsumerConfig(brokerCfg.Exchange.Name, "ans", "ans", "")
	if consConfig == nil {
		log.Fatal("Consumer config not created")
	}
	consConfig.AutoAck = true
	err = broker.CreateConsumer("ans_consumer", consConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	// читатель запросов
	consConfig = broker.NewConsumerConfig(brokerCfg.Exchange.Name, "ask", "ask", "")
	if consConfig == nil {
		log.Fatal("Consumer config not created")
	}

	err = broker.CreateConsumer("ask_consumer", consConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
