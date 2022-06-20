package slack

import (
	"fmt"

	"github.com/TKMAX777/ArkDiscordSlackLogNotifier/slack_webhook"
)

func (h *Handler) sendOnlineBlock() error {
	var OnlineMessageText = fmt.Sprintf("ArkOnlineMessage,%s", h.hook.Identity.UserID)
	var message = slack_webhook.Message{
		AsUser:   false,
		Channel:  h.settings.ChannelID,
		Username: h.settings.UserName,
		IconURL:  h.settings.AvaterURI,
		Blocks:   h.buildOnlineBlock(),
		Text:     OnlineMessageText,
	}

	var ts = h.lastMessageTS
	switch {
	case len(h.joinState.State) < 1:
		// there is no player
		if ts == "" {
			// if not found a last message, find from message history
			messages, err := h.hook.GetMessages(h.settings.ChannelID, "", 100)
			if err == nil {
				for _, msg := range messages {
					if msg.Text == OnlineMessageText {
						// *repost user messages contains DummyURIs
						ts = msg.TS
						break
					}
				}
			}

			if ts == "" {
				return nil
			}
		}
		h.lastMessageTS = ""
		h.hook.Remove(message.Channel, ts)
	default:
		// there are some players
		var ts = h.lastMessageTS
		if ts == "" {
			// if not found a last message, find from message history
			messages, err := h.hook.GetMessages(h.settings.ChannelID, "", 100)
			if err == nil {
				for _, msg := range messages {
					if msg.Text == OnlineMessageText {
						ts = msg.TS
						break
					}
				}
			}
		}
		if ts != "" {
			h.hook.Remove(message.Channel, ts)
		}

		message.TS = ts
		ts, err := h.hook.Send(message)
		if err != nil {
			return err
		}

		h.lastMessageTS = ts
	}

	return nil
}

func (h *Handler) buildOnlineBlock() []slack_webhook.BlockBase {
	var blocks = []slack_webhook.BlockBase{}

	var channelText = "Online"
	var channelNameElement = slack_webhook.MrkdwnElement(channelText)

	blocks = append(
		blocks,
		slack_webhook.ContextBlock(channelNameElement),
	)

	var userCount int
	var elements = []slack_webhook.BlockElement{}

	for username := range h.joinState.State {
		var userElm = slack_webhook.MrkdwnElement(h.settings.SendOptions.JoinAndLeftState.Emoji + " " + username)

		elements = append(elements, userElm)

		userCount++
		if userCount%4 == 0 {
			var block = slack_webhook.ContextBlock(elements...)
			blocks = append(blocks, block)

			elements = []slack_webhook.BlockElement{}
		}
	}

	if userCount%4 > 0 {
		var block = slack_webhook.ContextBlock(elements...)
		blocks = append(blocks, block)
	}

	blocks = append(blocks, slack_webhook.DividerBlock())

	return blocks
}
