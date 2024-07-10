package pubsub

type PubSub interface {
	Publish(topicId string, payload any) error
	// ValidatePayload(payload any) bool
	// Subscribe(topicId string, payload any) string
}
