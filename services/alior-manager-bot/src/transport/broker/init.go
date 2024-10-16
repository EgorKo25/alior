package broker

import (
	"alior-manager-bot/src/config"
	"github.com/EgorKo25/common/broker"
)

func NewBroker(cfg config.BrokerConfig) error {
	err := broker.InitBroker(cfg.URL)
	if err != nil {
		return err
	}

	pubConfig := broker.NewPublisherConfig(cfg.Exchange.Name, cfg.Exchange.Kind, cfg.Publisher.RoutingKey)
	if pubConfig == nil {
		return err
	}

	err = broker.CreatePublisher("ask_publisher", pubConfig)
	if err != nil {
		return err
	}

	consConfig := broker.NewConsumerConfig(cfg.Exchange.Name, cfg.Consumer.Name, cfg.Consumer.Queue, "")
	if consConfig == nil {
		return err
	}

	err = broker.CreateConsumer("ans_consumer", consConfig)
	if err != nil {
		return err
	}

	return nil
}
