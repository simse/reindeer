package server

import (
	"sort"
	"time"

	"github.com/mergestat/timediff"
	"github.com/simse/reindeer/internal/database"
	"github.com/simse/reindeer/internal/status"
)

type ServiceCategory struct {
	Name     string
	Services []map[string]interface{}
}

func StatusBannerInformation() map[string]string {
	// count failing services
	services := 0
	warn := 0
	fail := 0

	lastChecked := time.Time{}

	for _, service := range database.GetLatestServiceStatus() {
		services++

		if service.Status == status.ServiceStatusError {
			fail++
		}

		if service.Status == status.ServiceStatusWarn {
			warn++
		}

		if lastChecked.Before(service.LastChecked) {
			lastChecked = service.LastChecked
		}
	}

	// determine correct message
	cssClass := "ok"
	message := "All services are operational"

	if warn > 0 {
		cssClass = "warn"
		message = "Some services are having issues"

		if warn == services {
			message = "All services are having issues"
		}
	}

	if fail > 0 {
		cssClass = "error"
		message = "Some services are down"

		if fail == services {
			message = "All services are down"
		}
	}

	return map[string]string{
		"StatusStyle": cssClass,
		"Message":     message,
		"LastChecked": timediff.TimeDiff(lastChecked),
	}
}

var ServiceStyleMap = map[status.ServiceStatus]string{
	status.ServiceStatusError: "error",
	status.ServiceStatusWarn:  "warn",
	status.ServiceStatusOk:    "ok",
}
var ServiceTextMap = map[status.ServiceStatus]string{
	status.ServiceStatusError: "Down",
	status.ServiceStatusWarn:  "Partially Down",
	status.ServiceStatusOk:    "Operational",
}

func ServicesInformation() []ServiceCategory {
	// get status provider
	// TODO: get this somewhere else
	// provider := status.DummyStatusProvider{}

	// sort services by category
	temp := make(map[string][]status.Service)
	for _, service := range database.GetLatestServiceStatus() {
		temp[service.Category] = append(temp[service.Category], service)
	}

	// create and return ServiceCategory struct
	categories := []ServiceCategory{}

	for category, services := range temp {
		// convert services slice to displayable map
		displayableServices := []map[string]interface{}{}

		for _, service := range services {
			displayableServices = append(displayableServices, map[string]interface{}{
				"Name":     service.Name,
				"Status":   ServiceTextMap[service.Status],
				"CssClass": ServiceStyleMap[service.Status],
			})
		}

		categories = append(categories, ServiceCategory{
			Name:     category,
			Services: displayableServices,
		})
	}

	// sort by category a-z
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Name < categories[j].Name
	})

	return categories
}
