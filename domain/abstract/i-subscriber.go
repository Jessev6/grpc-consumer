package abstract

// ISubscriber is an interface that sends messages
type ISubscriber interface {
	Subscribe(ISubscription, *chan IMessage, *chan error)
}
