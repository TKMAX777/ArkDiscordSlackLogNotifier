package message_sender

import "github.com/TKMAX777/ArkDiscordSlackLogNotifier/ark"

type MessageSender func(ark.Message) error
