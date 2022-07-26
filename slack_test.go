package log

import (
	"os"
	"testing"

	"github.com/slack-go/slack"
)

func TestClient_SendWarn(t *testing.T) {
	client := slack.New(os.Getenv("SLACK_TOKEN"))
	logSlackClient := SlackClient{os.Getenv("SLACK_CHANNEL"), client}
	err := logSlackClient.SendWarn(
		"Something error occurred",
		"/example.v1.GreetService/Greet",
		"*error msg here*",
	)
	if err != nil {
		t.Fatalf("failed to send log: %v", err)
	}
}

func TestClient_SendError(t *testing.T) {
	client := slack.New(os.Getenv("SLACK_TOKEN"))
	logSlackClient := SlackClient{os.Getenv("SLACK_CHANNEL"), client}
	err := logSlackClient.SendError(
		"Something error occurred",
		"/example.v1.GreetService/Greet",
		"*error msg here*",
	)
	if err != nil {
		t.Fatalf("failed to send log: %v", err)
	}
}
