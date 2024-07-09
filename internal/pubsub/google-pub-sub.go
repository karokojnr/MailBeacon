package pubsub

import (
	gPubsub "cloud.google.com/go/pubsub"
)

type googlePubSub struct {
	pSub *gPubsub.Client
}

func NewGooglePubSub(pSub *gPubsub.Client) *googlePubSub {
	return &googlePubSub{
		pSub: pSub,
	}
}

func (g *googlePubSub) Publish(topiciD string, payload any) string {
	return ""
}

func (g *googlePubSub) Subscribe(topicId string, payload any) string {
	return ""
}
