# Golang Kafka: your Go app for easier Kafka server creation
    This is a simple Kafka integration within a Golang app. If you want to start a local Kafka server 
    to create, list and delete your topics, this is your place! 
    
    This repository is for testing and studying purposes, it's not meant to use as a production helper 
    on it's current state. We'll add more features within the time, 
    so follow us to be up to date with the new features =D

# USAGE

    This application runs with Docker Compose to create containers and initialize the Kafka server. 
    So, it is assumed that you have Docker and Docker Compose installed on your machine. 
    Beyond that, you'll need Golang installed too. 
    Knowing that, this is how you get to run the app:

    - Enter on the root folder and run docker-compose up -d;
    - Use the docker ps command to see if the Zookeeper and Kafka containers are running;


# COMMANDS

    - To create a topic: run "go run main.go --create-topic my_topic"
    - to list all topics: run "go run main.go --list-topics"

# CONTRIBUITION

Send an e-mail on "vinicius.mocci@hotmail.com" for further details on how to contribute with this repo.