package ark

type MessageTypeJoin struct {
	Content  string
	UserName string
}

type MessageTypeLeave struct {
	Content  string
	UserName string
}

type MessageTypeIlligalFormat struct{}

type MessageTypeOther struct {
	Content string
}
