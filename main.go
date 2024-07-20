package main

import (
	"github.com/djfemz/savannahTechTask/app/services"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file: ", err)
	}
}

func main() {
	//TODO: add channels to communicate
	go startJob()
}

func startJob() {
	job := cron.New()
	_, err := job.AddFunc("* */5 * * *", func() {
		services.FetchCommits()
	})
	if err != nil {
		log.Println("Error creating job: ", err)
	}
	log.Println("Starting new task...")
	job.Start()
}
