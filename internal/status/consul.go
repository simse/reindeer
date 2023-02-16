package status

import (
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

type ConsulStatusProvider struct{}

type ConsulService struct {
	Checks []struct {
		CheckID string `json:"CheckID"`
		Status  string `json:"Status"`
	} `json:"Checks"`
}

func (d *ConsulStatusProvider) GetServices() []Service {
	prefix := "reindeer"
	consulUrl := "http://100.72.238.121:8500"

	client := req.C().SetBaseURL(consulUrl)

	// get services in Consul catalog
	resp, err := client.R().Get("/v1/catalog/services")
	if err != nil {
		panic(err)
	}

	serviceCatalog := make(map[string][]string)
	err = resp.Into(&serviceCatalog)
	if err != nil {
		panic(err)
	}

	// check if monitoring is enabled by looking at tags
	servicesToCheck := make([]string, 0)

	for service, tags := range serviceCatalog {
		for _, tag := range tags {
			if tag == prefix+".enable=true" {
				servicesToCheck = append(servicesToCheck, service)
				break
			}
		}
	}

	services := make([]Service, 0)

	for _, service := range servicesToCheck {
		resp, err := client.R().Get("/v1/health/service/" + service)
		if err != nil {
			panic(err)
		}

		var foundServices []ConsulService
		err = resp.Into(&foundServices)
		if err != nil {
			panic(err)
		}

		if len(foundServices) == 0 {
			continue
		}

		// check if all checks are passing
		allChecksPassing := true

		for _, check := range foundServices[0].Checks {
			if check.Status != "passing" {
				allChecksPassing = false
				break
			}
		}

		status := ServiceStatusOk
		if !allChecksPassing {
			status = ServiceStatusError
		}

		// get service name from tags
		name := service

		// get description from tags
		description := ""

		// get category from tags
		category := ""

		for _, tag := range serviceCatalog[service] {
			if strings.HasPrefix(tag, prefix+".meta.title=") {
				name = strings.TrimPrefix(tag, prefix+".meta.title=")
			}

			if strings.HasPrefix(tag, prefix+".meta.description=") {
				description = strings.TrimPrefix(tag, prefix+".meta.description=")
			}

			if strings.HasPrefix(tag, prefix+".meta.category=") {
				category = strings.TrimPrefix(tag, prefix+".meta.category=")
			}
		}

		services = append(services, Service{
			Name:        name,
			Description: description,
			Category:    category,
			Status:      status,
			LastChecked: time.Now(),
			Provider:    "consul",
		})
	}

	return services
}
