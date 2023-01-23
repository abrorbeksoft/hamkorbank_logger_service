package events

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"logger_service/config"
	"logger_service/events/logger_service"
	"logger_service/pkg/logger"
	"logger_service/pkg/rabbitmq"
)

type PubSubServer struct {
	cfg      config.Config
	rabbitmq rabbitmq.RabbitMQI
	log      logger.LoggerI
}

func NewEvents(cfg config.Config, log logger.LoggerI, ch *amqp.Channel) (*PubSubServer, error) {
	rabbit, err := rabbitmq.NewRabbitMQ(cfg, ch)
	if err != nil {
		return nil, err
	}

	initPublishers(rabbit)

	return &PubSubServer{
		cfg:      cfg,
		log:      log,
		rabbitmq: rabbit,
	}, nil
}

func (s *PubSubServer) InitServices(ctx context.Context) {
	triggerListenerService := logger_service.NewTriggerListenerService(s.log, s.rabbitmq)
	triggerListenerService.RegisterConsumers()
	s.rabbitmq.RunConsumers(ctx)
}

func initPublishers(rabbit rabbitmq.RabbitMQI) {
	_ = rabbit.AddPublisher("v1.websocket_service.response")
}
