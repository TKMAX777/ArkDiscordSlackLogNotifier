package discord

import (
	"fmt"
	"strings"

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

			if h.settings.SendOptions.All.IsEnabled || h.settings.SendOptions.JoinAndLeft.IsEnabled {
				var text string
				switch {
				case onlineNumber > 1:
					text = fmt.Sprintf("%s %s\nOnline: %d players", h.settings.SendOptions.JoinAndLeft.Emoji, al.Content, onlineNumber)
				case onlineNumber == 1:
					text = fmt.Sprintf("%s %s\nOnline: %d player", h.settings.SendOptions.JoinAndLeft.Emoji, al.Content, onlineNumber)
				case onlineNumber <= 0:
					text = fmt.Sprintf("%s %s\nOnline: no players", h.settings.SendOptions.JoinAndLeft.Emoji, al.Content)
				}

				var message = discord_webhook.Message{
					Content:   strings.TrimSpace(text),
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

			if h.settings.SendOptions.All.IsEnabled || h.settings.SendOptions.JoinAndLeft.IsEnabled {
				var text string
				switch {
				case onlineNumber > 1:
					text = fmt.Sprintf("%s %s\nOnline: %d players", h.settings.SendOptions.JoinAndLeft.EmojiSub, al.Content, onlineNumber)
				case onlineNumber == 1:
					text = fmt.Sprintf("%s %s\nOnline: %d player", h.settings.SendOptions.JoinAndLeft.EmojiSub, al.Content, onlineNumber)
				case onlineNumber <= 0:
					text = fmt.Sprintf("%s %s\nOnline: no players", h.settings.SendOptions.JoinAndLeft.EmojiSub, al.Content)
				}

				var message = discord_webhook.Message{
					Content:   strings.TrimSpace(text),
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}

			return nil
		case ark.MessageTypeOther:
			if !h.settings.SendOptions.Other.IsEnabled {
				var message = discord_webhook.Message{
					Content:   strings.TrimSpace(fmt.Sprintf("%s %s", h.settings.SendOptions.Other.Emoji, al.Content)),
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeIlligalFormat:
			if !h.settings.SendOptions.All.IsEnabled {
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
