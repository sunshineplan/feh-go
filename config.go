package main

import (
	"encoding/json"
	"log"

	"github.com/sunshineplan/metadata"
	"github.com/sunshineplan/utils/mail"
)

var metadataConfig metadata.Config

func getMongo() (config mongoConfig) {
	m, err := metadataConfig.Get("feh_mongo")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(m, &config); err != nil {
		log.Fatal(err)
	}
	return
}

func getSubscribe() (config mail.Setting) {
	m, err := metadataConfig.Get("feh_subscribe")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(m, &config); err != nil {
		log.Fatal(err)
	}
	return
}

func getGithub() (config Github) {
	m, err := metadataConfig.Get("feh_github")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(m, &config); err != nil {
		log.Fatal(err)
	}
	return
}
