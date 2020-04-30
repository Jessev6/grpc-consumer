package abstract

// IMessage is the interface used for messages
type IMessage interface {
	Key() string
	ID() string
	Values() map[string]interface{}
}
