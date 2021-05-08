package main

import (
	"net/http"

	"github.com/jcuga/golongpoll"
)

func main() {
	// This uses the default/empty options. See section on customizing, and Options go docs.
	manager, err := golongpoll.StartLongpoll(golongpoll.Options{})
	if err != nil {
		panic(err)
	}
	// Expose pub-sub. You could omit the publish handler if you don't want
	// to allow clients to publish. For example, if clients only subscribe to data.
	http.HandleFunc("/events", manager.SubscriptionHandler)
	http.HandleFunc("/publish", manager.PublishHandler)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
