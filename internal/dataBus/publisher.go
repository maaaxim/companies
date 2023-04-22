package dataBus

import (
	"context"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"

	"github.com/any/companies/internal/infr/logger"
	"github.com/any/companies/internal/recovery"
	"github.com/any/companies/internal/services/common/events"
)

const (
	companiesPublishTopic = "companies.events"
)

type KafkaEventsPublisher struct {
	kafkaWriter *kafka.Writer
	logger      logger.Logger
	topic       string
}

func NewKafkaEventsPublisher(kafkaWriter *kafka.Writer, logger logger.Logger) *KafkaEventsPublisher {
	return &KafkaEventsPublisher{
		kafkaWriter: kafkaWriter,
		logger:      logger,
		topic:       companiesPublishTopic,
	}
}

func (p *KafkaEventsPublisher) GoPublishEvent(event events.Event) {
	go func() {
		defer recovery.RecoverToLog(p.logger)
		ctx := context.Background()
		p.sendMessage(ctx, p.makeMessage(event))
	}()
}

func (p *KafkaEventsPublisher) makeMessage(event events.Event) kafka.Message {
	messageJson, err := event.Marshal()
	if err != nil {
		p.logger.Error(errors.Wrapf(err, "kafka event marshal error. event name: %s", event.GetName()).Error())
	}

	return kafka.Message{
		Value: messageJson,
		Topic: p.topic,
	}
}

func (p *KafkaEventsPublisher) sendMessage(ctx context.Context, message kafka.Message) {
	switch err := p.kafkaWriter.WriteMessages(ctx, message).(type) { //nolint:errorlint
	case nil:
	case kafka.WriteErrors:
		if err[0] != nil {
			p.logger.Error(errors.Wrapf(err[0], "kafka send error. topic: %s", message.Topic).Error())
		}
	default:
		p.logger.Error(errors.Wrapf(err, "kafka send error. topic: %s", message.Topic).Error())
	}
}
