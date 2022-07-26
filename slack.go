package log

import (
	"github.com/slack-go/slack"
)

type SlackClient struct {
	Channel string
	*slack.Client
}

func (c *SlackClient) SendMsg(title string, source string, msg string) error {
	_, _, err := c.PostMessage(c.Channel, slack.MsgOptionBlocks(
		slack.NewHeaderBlock(slack.NewTextBlockObject(slack.PlainTextType, title, true, false)),
		slack.NewContextBlock("",
			slack.NewTextBlockObject(slack.PlainTextType, source, true, false),
		),
		slack.NewDividerBlock(),
		slack.NewContextBlock("",
			slack.NewTextBlockObject(slack.PlainTextType, msg, true, false),
		),
	))

	return err
}

func (c *SlackClient) SendWarn(title string, source string, msg string) error {
	title = "[⚠️] " + title
	return c.SendMsg(title, source, msg)
}

func (c *SlackClient) SendError(title string, source string, msg string) error {
	title = "[❗] " + title
	return c.SendMsg(title, source, msg)
}
