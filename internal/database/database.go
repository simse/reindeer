package database

import (
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/rs/zerolog/log"
	"github.com/simse/reindeer/internal/status"
)

var DB *storm.DB

func OpenDatabase(path string) {
	db, err := storm.Open(path)
	if err != nil {
		panic(err)
	}

	DB = db
}

func SaveServiceStatus(serviceStatus status.Service) {
	err := DB.Save(&serviceStatus)

	if err != nil {
		panic(err)
	}

	log.Info().Str("service_name", serviceStatus.Name).Msg("saved service status to database")
}

// get latest service status sorted by timestamp
func GetServiceStatus(serviceName string) status.Service {
	var services []status.Service

	err := DB.Select(q.Eq("Name", serviceName)).OrderBy("LastChecked").Reverse().Limit(1).Find(&services)
	if err != nil {
		panic(err)
	}

	return services[0]
}

// utility function to check if a string is in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// get all service names with no duplicates
func GetServiceNames() []string {
	var services []status.Service

	err := DB.All(&services)
	if err != nil {
		panic(err)
	}

	var serviceNames []string
	for _, service := range services {
		if !contains(serviceNames, service.Name) {
			serviceNames = append(serviceNames, service.Name)
		}
	}

	return serviceNames
}

// get latest service status for all services
func GetLatestServiceStatus() []status.Service {
	var services []status.Service

	for _, serviceName := range GetServiceNames() {
		services = append(services, GetServiceStatus(serviceName))
	}

	// remove services whose LastChecked is more than 2 minutes ago
	var returnServices []status.Service
	for _, service := range services {
		if service.LastChecked.After(time.Now().Add(-2 * time.Minute)) {
			returnServices = append(returnServices, service)
		}
	}

	return returnServices
}
