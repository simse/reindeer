package main

import (
	"time"

	"github.com/onatm/clockwerk"
	"github.com/simse/reindeer/internal/config"
	"github.com/simse/reindeer/internal/database"
	"github.com/simse/reindeer/internal/server"
)

func init() {
	database.OpenDatabase("/home/simon/projects/reindeer/reindeer.db")

	config.LoadConfig()
}

func main() {
	var job server.GetServicesJob

	c := clockwerk.New()
	c.Every(60 * time.Second).Do(job)
	c.Start()

	go job.Run()

	server.StartServer()
}
