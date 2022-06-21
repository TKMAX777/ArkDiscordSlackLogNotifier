package ark

import "time"

type Message struct {
	RawLine string
	Content string
	Time    time.Time
	Event   interface{}
}

type MessageTypeJoin struct {
	User User
}

type MessageTypeLeave struct {
	User User
}

type MessageTypeKilled struct {
	KilledUser User
	KilledBy   Enemy
}

type MessageTypeIlligalFormat struct{}

type MessageTypeOther struct{}
