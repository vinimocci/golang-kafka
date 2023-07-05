package main

import (
	"os"
	"vinimocci/golang-kafka/server"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		panic("missing at least one argument")
	}

	switch args[1] {
    case "--create-topic":
		topicToCreate := args[2]
        server.CreateTopicByGivenName(topicToCreate)
    case "--list-topics":
		server.ListAllTopics()
    default:
        panic("missing a valid argument")
    }
}