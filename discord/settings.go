package discord

type Settings struct {
	Token       string
	HookURI     string
	SendOptions SendOptions

	ChannelID string
	AvaterURI string
	UserName  string
}

type SendOptions struct {
	JoinAndLeftState SendOption
	JoinAndLeft      SendOption
	Other            SendOption
	All              SendOption
}

type SendOption struct {
	IsEnabled bool
	Emoji     string
	EmojiSub  string
}
