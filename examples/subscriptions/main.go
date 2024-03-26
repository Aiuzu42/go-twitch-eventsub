package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	twitcheventsub "github.com/Aiuzu42/go-twitch-eventsub"
)

func main() {
	token := ""
	clientID := ""
	secret := ""
	broadcasterId := ""
	callback := ""
	crtPath := ""
	keyPath := ""
	fmt.Println("Starting...")

	//Create a new instance of the Client and define functions for errors and an event, in this example channel.follow
	client := twitcheventsub.NewClient(crtPath, keyPath, secret, callback)
	client.OnError(handleError)
	client.OnChannelFollow(handleFollow)

	//We subscribe to the channel.follow event
	sub, err := client.SubscribeToEvent(twitcheventsub.Follow, broadcasterId, token, clientID)
	if err != nil {
		panic("unable to create subscription: " + err.Error())
	} else {
		fmt.Println("Subscription created: " + sub.Data[0].Id)
	}

	go client.StartServer()

	//Wait for an os.Signal to delete example subscription and stop server
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	//Delete the subscription created for the example
	err = client.DeleteSubscription(sub.Data[0].Id, token, clientID)
	if err != nil {
		fmt.Println("unable to delete subscription " + sub.Data[0].Id + " " + err.Error())
	}

	//Server stop
	err = client.StopServer()
	if err != nil {
		fmt.Println("unable to stop server: " + err.Error())
	}
	fmt.Println("Closing!")
}

func handleError(err error) {
	fmt.Println(err)
}

func handleFollow(event twitcheventsub.ChannelFollowEvent) {
	fmt.Println("Follow from: " + event.UserName)
}
