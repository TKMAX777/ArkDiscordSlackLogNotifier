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
	return message_sender.MessageSender(func(arklog ark.Message) error {

		switch al := arklog.Event.(type) {
		case ark.MessageTypeJoin:
			h.joinState.Join(al.User.Name)

			// if h.settings.SendOptions.JoinAndLeftState {
			// 		TODO: Make State Image
			// }

			if h.settings.SendOptions.All.IsEnabled || h.settings.SendOptions.JoinAndLeft.IsEnabled {
				var text string
				switch {
				case len(h.joinState.State) > 1:
					text = fmt.Sprintf("%s %s\nOnline: %d players", h.settings.SendOptions.JoinAndLeft.Emoji, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) == 1:
					text = fmt.Sprintf("%s %s\nOnline: %d player", h.settings.SendOptions.JoinAndLeft.Emoji, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) <= 0:
					text = fmt.Sprintf("%s %s\nOnline: no players", h.settings.SendOptions.JoinAndLeft.Emoji, arklog.Content)
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
			h.joinState.Leave(al.User.Name)

			// if h.settings.SendOptions.JoinAndLeftState {
			// 		TODO: Make State Image
			// }

			if h.settings.SendOptions.All.IsEnabled || h.settings.SendOptions.JoinAndLeft.IsEnabled {
				var text string
				switch {
				case len(h.joinState.State) > 1:
					text = fmt.Sprintf("%s %s\nOnline: %d players", h.settings.SendOptions.JoinAndLeft.EmojiSub, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) == 1:
					text = fmt.Sprintf("%s %s\nOnline: %d player", h.settings.SendOptions.JoinAndLeft.EmojiSub, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) <= 0:
					text = fmt.Sprintf("%s %s\nOnline: no players", h.settings.SendOptions.JoinAndLeft.EmojiSub, arklog.Content)
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
		case ark.MessageTypeTamed:
			if h.settings.SendOptions.Tamed.IsEnabled {
				var message = discord_webhook.Message{
					Content:   strings.TrimSpace(fmt.Sprintf("%s %s", h.settings.SendOptions.Tamed.Emoji, arklog.Content)),
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeKilled:
			if h.settings.SendOptions.Killed.IsEnabled {
				var message = discord_webhook.Message{
					Content:   strings.TrimSpace(fmt.Sprintf("%s %s", h.settings.SendOptions.Killed.Emoji, arklog.Content)),
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeOther:
			if h.settings.SendOptions.Other.IsEnabled {
				var message = discord_webhook.Message{
					Content:   strings.TrimSpace(fmt.Sprintf("%s %s", h.settings.SendOptions.Other.Emoji, arklog.Content)),
					ChannelID: h.settings.ChannelID,
					UserName:  h.settings.UserName,
					AvaterURL: h.settings.AvaterURI,
				}

				_, err := h.hook.Send(h.settings.ChannelID, message, false, nil)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeIlligalFormat:
			if h.settings.SendOptions.All.IsEnabled {
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
