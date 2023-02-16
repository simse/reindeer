package server

import (
	"github.com/rs/zerolog/log"
	"github.com/simse/reindeer/internal/database"
	"github.com/simse/reindeer/internal/status"
)

type GetServicesJob struct{}

func (j GetServicesJob) Run() {
	// get status provider
	provider := status.ConsulStatusProvider{}

	log.Info().Msg("getting service status from provider")

	for _, service := range provider.GetServices() {
		database.SaveServiceStatus(service)
	}
}
