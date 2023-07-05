package server

import (
	"fmt"
	"vinimocci/golang-kafka/entities/topics"
)

func CreateTopicByGivenName(topicToCreate string){
	fmt.Printf("Creating topic: %s\n", topicToCreate)

	_, generationErr := topics.CreateTopic(topicToCreate)
	if generationErr != nil{
		panic(generationErr)
	}

	fmt.Println("Topic created successfully!")
}

func ListAllTopics(){
	allTopics, allTpErr := topics.ListAllTopics()
	if allTpErr != nil{
		panic(allTpErr)
	}

	fmt.Println("showing topics: ")
	
	for topic := range allTopics {
		fmt.Println(topic)
	}
}

func DeleteTopic(topicToDelete string){
	fmt.Printf("Deleting topic: %s\n", topicToDelete)

	_, deleteError := topics.DeleteTopic(topicToDelete)
	if deleteError != nil{
		panic(deleteError)
	}

	fmt.Println("Topic deleted successfully!")
}