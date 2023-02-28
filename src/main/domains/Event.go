package main_domains

type Event struct {
	Key   []byte
	Value []byte
}

func NewEvent(key string, value string) Event {
	return Event{
		Key:   []byte(key),
		Value: []byte(value),
	}
}
