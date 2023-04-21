package kafka

import (
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type WriterConfig struct {
	Hosts                    string `envconfig:"HOSTS"`
	BatchTimeoutMilliseconds int    `envconfig:"BATCH_TIMEOUT_MILLISECONDS" default:"500"`
	MaxAttempts              int    `envconfig:"MAX_ATTEMPTS" default:"10"`
	RequiredAcks             string `envconfig:"REQUIRED_ACKS" default:"all"`
	IsAutoTopicCreation      int    `envconfig:"IS_AUTO_TOPIC_CREATION" default:"0"`
}

func InitWriter(cfg WriterConfig) (*kafka.Writer, error) {
	var requiredAcks kafka.RequiredAcks
	err := requiredAcks.UnmarshalText([]byte(cfg.RequiredAcks))
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	isAutoTopicCreation := false
	if cfg.IsAutoTopicCreation > 0 {
		isAutoTopicCreation = true
	}

	return &kafka.Writer{
		Addr:                   kafka.TCP(strings.Split(cfg.Hosts, ",")...),
		Compression:            compress.Lz4,
		BatchTimeout:           time.Duration(cfg.BatchTimeoutMilliseconds) * time.Millisecond,
		MaxAttempts:            cfg.MaxAttempts,
		RequiredAcks:           requiredAcks,
		AllowAutoTopicCreation: isAutoTopicCreation,
	}, nil
}
