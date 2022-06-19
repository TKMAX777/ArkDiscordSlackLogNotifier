package slack

import (
	"fmt"

	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/ark"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/message_sender"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/slack_webhook"
	"github.com/pkg/errors"
)

func (h *Handler) SendMessageFunction() message_sender.MessageSender {
	var onlineNumber int

	return message_sender.MessageSender(func(arklog ark.Message) error {

		switch al := arklog.Content.(type) {
		case ark.MessageTypeJoin:
			onlineNumber += 1

			h.joinState.Join(al.UserName)

			if h.settings.SendOptions.JoinAndLeftState {
				err := h.sendOnlineBlock()
				if err != nil {
					return errors.Wrap(err, "sendOnlineBlock")
				}
			}

			if h.settings.SendOptions.All || h.settings.SendOptions.JoinAndLeft {
				var text string
				switch {
				case onlineNumber > 1:
					text = fmt.Sprintf("%s\nOnline: %d players", al.Content, onlineNumber)
				case onlineNumber == 1:
					text = fmt.Sprintf("%s\nOnline: %d player", al.Content, onlineNumber)
				case onlineNumber <= 0:
					text = fmt.Sprintf("%s\nOnline: no players", al.Content)
				}

				var message = slack_webhook.Message{
					Text:     text,
					Channel:  h.settings.ChannelID,
					Username: h.settings.UserName,
					IconURL:  h.settings.AvaterURI,
					AsUser:   false,
				}

				_, err := h.hook.Send(message)

				return errors.Wrap(err, "Send")
			}

			return nil
		case ark.MessageTypeLeave:
			onlineNumber -= 1

			h.joinState.Leave(al.UserName)

			if h.settings.SendOptions.JoinAndLeftState {
				err := h.sendOnlineBlock()
				if err != nil {
					return errors.Wrap(err, "sendOnlineBlock")
				}
			}

			if h.settings.SendOptions.All || h.settings.SendOptions.JoinAndLeft {
				var text string
				switch {
				case onlineNumber > 1:
					text = fmt.Sprintf("%s\nOnline: %d players", al.Content, onlineNumber)
				case onlineNumber == 1:
					text = fmt.Sprintf("%s\nOnline: %d player", al.Content, onlineNumber)
				case onlineNumber <= 0:
					text = fmt.Sprintf("%s\nOnline: no players", al.Content)
				}

				var message = slack_webhook.Message{
					Text:     text,
					Channel:  h.settings.ChannelID,
					Username: h.settings.UserName,
					IconURL:  h.settings.AvaterURI,
					AsUser:   false,
				}

				_, err := h.hook.Send(message)

				return errors.Wrap(err, "Send")
			}

			return nil
		case ark.MessageTypeOther:
			if !h.settings.SendOptions.Other {
				var message = slack_webhook.Message{
					Text:     al.Content,
					Channel:  h.settings.ChannelID,
					Username: h.settings.UserName,
					IconURL:  h.settings.AvaterURI,
					AsUser:   false,
				}

				_, err := h.hook.Send(message)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeIlligalFormat:
			if !h.settings.SendOptions.All {
				return nil
			}

			var message = slack_webhook.Message{
				Text:     arklog.RawLine,
				Channel:  h.settings.ChannelID,
				Username: h.settings.UserName,
				IconURL:  h.settings.AvaterURI,
				AsUser:   false,
			}

			_, err := h.hook.Send(message)

			return errors.Wrap(err, "Send")
		}
		return nil
	})
}
