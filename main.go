package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func goGetEnvVar(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	/*botToken := goGetEnvVar("SLACK_BOT_TOKEN")
	appToken := goGetEnvVar("SLACK_APP_TOKEN")*/

	bot := slacker.NewClient("xoxb-4010050576038-4010120797846-25hEc7SQzeG69op8arWRDRNC", "xapp-1-A040V8Q8GBB-4010072985622-857c741643ce25f9c8374ddb030d0c1f8cef544d17055af8967a33c7030f30b1")

	go printCommandEvents(bot.CommandEvents())

	var examples []string

	example := append(examples, "My YOB is 2020")

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB calculator",
		Examples:    example,
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println(err)
			}
			age := 2022 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}
