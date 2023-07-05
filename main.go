package main

import (
	"os"
	"vinimocci/golang-kafka/server"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		panic("missing at least one argument with it's value")
	}

	givenValue := args[2]

	switch args[1] {
    case "--create-topic":
        server.CreateTopicByGivenName(givenValue)

    case "--list-topics":
		server.ListAllTopics()

	case "--delete-topic":
		server.DeleteTopic(givenValue)

    default:
        panic("missing a valid argument")
    }
}