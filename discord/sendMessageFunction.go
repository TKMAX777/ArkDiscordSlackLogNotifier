package discord

import (
	"fmt"

	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/ark"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/discord_webhook"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/message_sender"
	"github.com/pkg/errors"
)

func (h *Handler) SendMessageFunction() message_sender.MessageSender {
	var onlineNumber int

	return message_sender.MessageSender(func(arklog ark.Message) error {

		switch al := arklog.Content.(type) {
		case ark.MessageTypeJoin:
			onlineNumber += 1

			h.joinState.Join(al.UserName)

			// if h.settings.SendOptions.JoinAndLeftState {
			// 		TODO: Make State Image
			// }

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

				var message = discord_webhook.Message{
					Content:   text,
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}

			return nil
		case ark.MessageTypeLeave:
			onlineNumber -= 1

			h.joinState.Leave(al.UserName)

			// if h.settings.SendOptions.JoinAndLeftState {
			// 		TODO: Make State Image
			// }

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

				var message = discord_webhook.Message{
					Content:   text,
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}

			return nil
		case ark.MessageTypeOther:
			if !h.settings.SendOptions.Other {
				var message = discord_webhook.Message{
					Content:   al.Content,
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeIlligalFormat:
			if !h.settings.SendOptions.All {
				return nil
			}

			var message = discord_webhook.Message{
				Content:   arklog.RawLine,
				ChannelID: h.settings.ChannelID,
				UserName:  h.settings.UserName,
				AvaterURL: h.settings.AvaterURI,
			}

			_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

			return errors.Wrap(err, "Send")
		}
		return nil
	})
}
