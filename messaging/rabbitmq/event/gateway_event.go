package event

type MessageDeliverEvent struct {
	Recipients []string
	Body       []byte
	Protocol   string
}

type ConnCreateEvent struct {
	Id      string
	Address string
}

type ConnCreateCmd struct {
	Address string
}