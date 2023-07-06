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

	switch args[1] {
    case "--create-topic":
		givenValue := args[2]
        server.CreateTopicByGivenName(givenValue)

    case "--list-topics":
		server.ListAllTopics()

	case "--delete-topic":
		givenValue := args[2]
		server.DeleteTopic(givenValue)
	
	case "--create-message":
		server.CreateStringMessageOnGivenTopic(args[4], args[2])

    default:
        panic("missing a valid argument")
    }
}