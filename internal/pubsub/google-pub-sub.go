package pubsub

import (
	"context"
	"encoding/json"
	"os"

	gPubsub "cloud.google.com/go/pubsub"
	_ "github.com/joho/godotenv/autoload"
)

var (
	developmentEnv = os.Getenv("DEVELOPMENT_ENV")
)

type googlePubSub struct {
	pSub *gPubsub.Client
}

func NewGooglePubSub(pSub *gPubsub.Client) *googlePubSub {
	return &googlePubSub{
		pSub: pSub,
	}
}

func (g *googlePubSub) Publish(topicId string, payload any) error {

	topic := g.pSub.Topic(topicId + "-" + developmentEnv)
	if ok, err := topic.Exists(context.Background()); err != nil || !ok {
		return err
	}

	p, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	res := topic.Publish(context.Background(), &gPubsub.Message{
		Data: p,
	})

	_, err = res.Get(context.Background())
	if err != nil {
		return err
	}

	return nil
}
