package abstract

// ISubscription is an interface for subscriptions
type ISubscription interface {
	Key() string
	From() string
	To() string
	Count() int32
	IncludeMetadata() bool
}
