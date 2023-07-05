package topics

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func CreateTopic(topicName string) (bool, error) {
	brokers := []string{"localhost:9092"}

	admin, admErr := sarama.NewClusterAdmin(brokers, nil)
	if admErr != nil {
		return false, admErr
	}

	topicDetails := sarama.TopicDetail{
		NumPartitions: 1,
		ReplicationFactor: 1,	
	}

	createErr := admin.CreateTopic(topicName, &topicDetails, false)
	if createErr != nil{
		return false, createErr
	}

	createErr = admin.Close()
	if createErr != nil{
		return false, nil
	}

	return true, nil
}

func ListAllTopics() (map[string]sarama.TopicDetail, error) {
	brokers := []string{"localhost:9092"}

	admin, admErr := sarama.NewClusterAdmin(brokers, nil)
	if admErr != nil {
		return nil, admErr
	}

	topics, err := admin.ListTopics()
	if err != nil {
		return nil, err
	}

	if len(topics) == 0 {
		return nil, fmt.Errorf("no topics to show") 
	}

	return topics, nil
}

func DeleteTopic(topicName string) (bool, error) {
	brokers := []string{"localhost:9092"}

	config := sarama.NewConfig()
	admin, adminErr := sarama.NewClusterAdmin(brokers, config)
	if adminErr != nil {
		return false, adminErr
	}
	defer admin.Close()


	deleteErr := admin.DeleteTopic(topicName)
	if deleteErr != nil {
		return false, deleteErr
	}

	return true, nil
}