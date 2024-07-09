package pubsub

type googlePubSub struct {
	projectId string
}

type GooglePubSub interface {
	Publish(topiciD string, payload any) string
	Subscribe(topicId string, payload any) string
}

func NewGooglePubSub(projectId string) *googlePubSub {
	return &googlePubSub{
		projectId: projectId,
	}
}
