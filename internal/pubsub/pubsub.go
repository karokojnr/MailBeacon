package pubsub

type PubSub interface {
	Publish(topiciD string, payload any) string
	Subscribe(topicId string, payload any) string
}
