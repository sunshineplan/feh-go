package main

import (
	"log"

	"github.com/sunshineplan/utils/mail"
	"github.com/sunshineplan/utils/metadata"
)

var meta metadata.Server

func getMongo() (config mongoConfig) {
	if err := meta.Get("feh_mongo", &config); err != nil {
		log.Fatal(err)
	}
	return
}

func getSubscribe() (dialer *mail.Dialer, to []string) {
	var config struct {
		SMTPServer     string
		SMTPServerPort int
		From, Password string
		To             []string
	}
	if err := meta.Get("feh_subscribe", &config); err != nil {
		log.Fatalln("Failed to get feh_subscribe metadata:", err)
	}
	dialer.Host = config.SMTPServer
	dialer.Port = config.SMTPServerPort
	dialer.Account = config.From
	dialer.Password = config.Password
	to = config.To
	return
}

func getGithub() (config Github) {
	if err := meta.Get("feh_github", &config); err != nil {
		log.Fatal(err)
	}
	return
}