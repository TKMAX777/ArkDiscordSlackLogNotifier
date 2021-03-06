package slack

type Settings struct {
	Token       string
	SendOptions SendOptions

	ChannelID string
	AvaterURI string
	UserName  string
}

type SendOptions struct {
	JoinAndLeftState SendOption
	JoinAndLeft      SendOption
	Tamed            SendOption
	Killed           SendOption
	Other            SendOption
	All              SendOption
}

type SendOption struct {
	IsEnabled bool
	Emoji     string
	EmojiSub  string
}
