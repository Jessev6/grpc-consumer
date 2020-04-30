package domain

// EventSubscription is used to subscribe for events
type EventSubscription struct {
	key             string
	from            string
	to              string
	count           int32
	includeMetadata bool
	reversed        bool
}

// NewEventSubscription is the constructor for EventSubscriptions
func NewEventSubscription(key string, from string, to string, count int32, includeMetadata bool) *EventSubscription {
	reversed := false
	if from == "" {
		from = "-"
	}

	if to == "" {
		to = "+"
	}

	if count < 0 {
		if from == "-" {
			from = "+"
		}

		if to == "+" {
			to = "-"
		}

		count *= -1
		reversed = true
	}

	return &EventSubscription{
		key:             key,
		from:            from,
		to:              to,
		count:           count,
		includeMetadata: includeMetadata,
		reversed:        reversed,
	}
}

// Key is the getter for property key
func (s *EventSubscription) Key() string {
	return s.key
}

// From is the getter for from
func (s *EventSubscription) From() string {
	return s.from
}

// To is the getter for property to
func (s *EventSubscription) To() string {
	return s.to
}

// Count is the getter for property count
func (s *EventSubscription) Count() int32 {
	return s.count
}

// IncludeMetadata is the getter for property includeMetadata
func (s *EventSubscription) IncludeMetadata() bool {
	return s.includeMetadata
}

// Reversed is the getter for property reversed
func (s *EventSubscription) Reversed() bool {
	return s.reversed
}
