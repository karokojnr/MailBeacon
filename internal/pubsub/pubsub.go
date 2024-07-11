package pubsub

type PubSub interface {
	Publish(topicId string, payload any) error
}
