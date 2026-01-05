package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
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
	client := twitcheventsub.NewClient(secret, callback)
	client.OnError(handleError)
	client.OnChannelFollow(handleFollow)

	//We subscribe to the channel.follow event
	sub, err := client.SubscribeToEvent(twitcheventsub.Follow, broadcasterId, token, clientID)
	if err != nil {
		panic("unable to create subscription: " + err.Error())
	} else {
		fmt.Println("Subscription created: " + sub.Data[0].Id)
	}

	//Setup the tls server to listen for incoming events with the handler function
	srv := &http.Server{Addr: ":443"}
	stop := &sync.WaitGroup{}
	stop.Add(1)
	http.HandleFunc("/eventsub", client.HandleEvent)
	defer stop.Done()
	go func() {
		if err := srv.ListenAndServeTLS(crtPath, keyPath); err != http.ErrServerClosed {
			panic("eventsub start server: " + err.Error())
		}
	}()

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
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("unable to stop server: " + err.Error())
	}
	stop.Wait()
	fmt.Println("Closing!")
}

func handleError(err error) {
	fmt.Println(err)
}

func handleFollow(event twitcheventsub.ChannelFollowEvent) {
	fmt.Println("Follow from: " + event.UserName)
}
