package main

import (
	"context"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-io/slacker"
)

// Showcasing the ability to add attachments to a `Reply`

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	definition := &slacker.CommandDefinition{
		Command:     "echo {word}",
		Description: "Echo a word!",
		Handler: func(ctx *slacker.CommandContext) {
			word := ctx.Request().Param("word")

			attachments := []slack.Attachment{}
			attachments = append(attachments, slack.Attachment{
				Color:      "good",
				AuthorName: "Raed Shomali",
				Title:      "Attachment Title",
				Text:       "Attachment Text",
			})

			ctx.Response().Reply(word, slacker.WithAttachments(attachments))
		},
	}

	bot.AddCommand(definition)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
