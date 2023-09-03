package nats

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"time"
	"wb_l0/config"
	"wb_l0/internal/models"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type Nats struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
	sc     stan.Conn
	ns     *nats.Conn
}

func NewNats(cfg *config.Config, logger *zap.SugaredLogger) *Nats {

	url := fmt.Sprintf("nats://%s:%s", cfg.Nats.Host, cfg.Nats.Port)
	ns, err := nats.Connect(url)
	if err != nil {
		logger.Error(err)
		return nil
	}

	sc, err := stan.Connect(cfg.Nats.Cluster, cfg.Nats.Client)
	if err != nil {
		logger.Error(err)
		return nil
	}

	return &Nats{cfg: cfg, logger: logger,
		sc: sc, ns: ns}
}

func (n *Nats) PublishMessage(subject string, message models.OrderModel) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return n.sc.Publish(subject, messageJSON)
}

func (n *Nats) SubscribeAndReceiveMessage(subject string) (*models.OrderModel, error) {

	var receivedMessage models.OrderModel

	ch := make(chan *models.OrderModel)
	_, err := n.sc.Subscribe(subject, func(msg *stan.Msg) {
		//var receivedMessage models.OrderModel
		err := json.Unmarshal(msg.Data, &receivedMessage)
		if err != nil {
			n.logger.Errorf("Error unmarshalling message:", err)
			return
		}
		ch <- &receivedMessage
	})
	if err != nil {
		n.logger.Errorf("Error in subscribing on subject:", err)
		return nil, err
	}

	// Ждем получения сообщения в течение 5 секунд (может потребоваться настройка)
	select {
	case receivedMessage := <-ch:
		return receivedMessage, nil
	case <-time.After(60 * time.Second):
		return nil, stan.ErrTimeout
	}
}
