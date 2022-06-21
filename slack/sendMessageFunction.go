package slack

import (
	"fmt"
	"strings"

	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/ark"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/message_sender"
	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/slack_webhook"
	"github.com/pkg/errors"
)

func (h *Handler) SendMessageFunction() message_sender.MessageSender {
	return message_sender.MessageSender(func(arklog ark.Message) error {

		switch al := arklog.Event.(type) {
		case ark.MessageTypeJoin:
			h.joinState.Join(al.User.Name)

			if h.settings.SendOptions.JoinAndLeftState.IsEnabled {
				err := h.sendOnlineBlock()
				if err != nil {
					return errors.Wrap(err, "sendOnlineBlock")
				}
			}

			if h.settings.SendOptions.All.IsEnabled || h.settings.SendOptions.JoinAndLeft.IsEnabled {
				var text string
				switch {
				case len(h.joinState.State) > 1:
					text = fmt.Sprintf("%s %s\nOnline: %d players", h.settings.SendOptions.JoinAndLeftState.Emoji, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) == 1:
					text = fmt.Sprintf("%s %s\nOnline: %d player", h.settings.SendOptions.JoinAndLeftState.Emoji, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) <= 0:
					text = fmt.Sprintf("%s %s\nOnline: no players", h.settings.SendOptions.JoinAndLeftState.Emoji, arklog.Content)
				}

				var message = slack_webhook.Message{
					Text:     strings.TrimSpace(text),
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
			h.joinState.Leave(al.User.Name)

			if h.settings.SendOptions.JoinAndLeftState.IsEnabled {
				err := h.sendOnlineBlock()
				if err != nil {
					return errors.Wrap(err, "sendOnlineBlock")
				}
			}

			if h.settings.SendOptions.All.IsEnabled || h.settings.SendOptions.JoinAndLeft.IsEnabled {
				var text string
				switch {
				case len(h.joinState.State) > 1:
					text = fmt.Sprintf("%s %s\nOnline: %d players", h.settings.SendOptions.JoinAndLeftState.EmojiSub, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) == 1:
					text = fmt.Sprintf("%s %s\nOnline: %d player", h.settings.SendOptions.JoinAndLeftState.EmojiSub, arklog.Content, len(h.joinState.State))
				case len(h.joinState.State) <= 0:
					text = fmt.Sprintf("%s %s\nOnline: no players", h.settings.SendOptions.JoinAndLeftState.EmojiSub, arklog.Content)
				}

				var message = slack_webhook.Message{
					Text:     strings.TrimSpace(text),
					Channel:  h.settings.ChannelID,
					Username: h.settings.UserName,
					IconURL:  h.settings.AvaterURI,
					AsUser:   false,
				}

				_, err := h.hook.Send(message)

				return errors.Wrap(err, "Send")
			}

			return nil
		case ark.MessageTypeKilled:
			if h.settings.SendOptions.Other.IsEnabled {
				var message = slack_webhook.Message{
					Text:     strings.TrimSpace(fmt.Sprintf("%s %s", h.settings.SendOptions.Killed.Emoji, arklog.Content)),
					Channel:  h.settings.ChannelID,
					Username: h.settings.UserName,
					IconURL:  h.settings.AvaterURI,
					AsUser:   false,
				}

				_, err := h.hook.Send(message)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeOther:
			if h.settings.SendOptions.Other.IsEnabled {
				var message = slack_webhook.Message{
					Text:     strings.TrimSpace(fmt.Sprintf("%s %s", h.settings.SendOptions.Other.Emoji, arklog.Content)),
					Channel:  h.settings.ChannelID,
					Username: h.settings.UserName,
					IconURL:  h.settings.AvaterURI,
					AsUser:   false,
				}

				_, err := h.hook.Send(message)

				return errors.Wrap(err, "Send")
			}
		case ark.MessageTypeIlligalFormat:
			if !h.settings.SendOptions.All.IsEnabled {
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
