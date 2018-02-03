package comms

// Message is a simple string message that is received from the web to be added to the chain
type Message string

func (m Message) String() string {
	return string(m)
}
