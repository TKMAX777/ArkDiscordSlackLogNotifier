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
	JoinAndLeftState bool
	JoinAndLeft      bool
	Other            bool
	All              bool
}
