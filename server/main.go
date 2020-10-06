package main

import (
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
	msg := &i18n.Message{
		ID:    "poll.updateVote.maxVotes",
		One:   "one",
		Few:   "few",
		Many:  "many",
		Other: "other",
	}
	log.Printf("%#v", msg)
}
