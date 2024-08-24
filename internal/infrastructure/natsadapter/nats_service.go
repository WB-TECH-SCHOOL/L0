package natsadapter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"l0/internal/cache"
	"l0/internal/errs"
	"l0/internal/infrastructure/config"
	"l0/internal/models"
	"l0/internal/repository"
	"time"
)

const (
	messages = "messages."

	messagesRoot = messages + "*"
)

type NatsService struct {
	js             nats.JetStreamContext
	orderRepo      repository.Orders
	dbResponseTime time.Duration
	cache          cache.Cache
	logger         *zerolog.Logger
}

func InitNATSService(
	db *sqlx.DB,
	cache cache.Cache,
	logger *zerolog.Logger,
) *NatsService {
	connectUrl := fmt.Sprintf("nats://%s:%s", viper.GetString(config.NATSHost), viper.GetString(config.NATSPort))
	conn, err := nats.Connect(connectUrl)
	if err != nil {
		panic(err)
	}

	js, err := conn.JetStream()
	if err != nil {
		panic(err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "MESSAGES",
		Subjects: []string{messagesRoot},
		Storage:  nats.FileStorage,
	})
	if err != nil {
		panic(err)
	}

	return &NatsService{
		js:             js,
		orderRepo:      repository.InitOrderRepository(db),
		dbResponseTime: time.Duration(viper.GetInt(config.DBResponseTime)) * time.Second,
		cache:          cache,
		logger:         logger,
	}
}

func (n *NatsService) Listen() {
	_, err := n.js.Subscribe(messagesRoot, func(msg *nats.Msg) {
		n.logger.Info().Msgf("Received message: %s from %s\n", string(msg.Data), msg.Subject)

		meta, err := msg.Metadata()
		if err != nil {
			n.logger.Error().Err(err).Msg(errs.ErrNatsMetadata.Error())
			msg.Nak()
			return
		}

		receivedData := make(map[string]any)
		if err := json.Unmarshal(msg.Data, &receivedData); err != nil {
			n.logger.Error().Err(err).Msg(errs.ErrUnmarshal.Error())
			if meta.NumDelivered > 3 {
				msg.Term()
			} else {
				msg.Nak()
			}
			return
		}

		if err = validateData(receivedData); err != nil {
			n.logger.Error().Err(err)
			msg.Term()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), n.dbResponseTime)
		defer cancel()
		order := models.Order{
			ID:   receivedData["order_uid"].(string),
			Data: msg.Data,
		}
		if err := n.orderRepo.Create(ctx, order); err != nil {
			n.logger.Error().Err(err)
			msg.Nak()
			return
		}

		n.cache.Add(order)
		msg.Ack()
	})
	if err != nil {
		panic(err)
	}
}

func validateData(data map[string]any) error {
	if data["order_uid"] == nil {
		return errs.ErrNatsInvalidData
	}
	return nil
}
