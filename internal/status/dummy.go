package status

import "time"

type DummyStatusProvider struct{}

func (d *DummyStatusProvider) GetServices() []Service {
	return []Service{
		{
			Name:        "Nomad",
			Description: "Workload orchestrator service",
			Category:    "Infrastructure",
			Status:      ServiceStatusOk,
			LastChecked: time.Now(),
			Provider:    "dummy",
		},
		{
			Name:        "Waypoint",
			Description: "Application service",
			Category:    "Infrastructure",
			Status:      ServiceStatusOk,
			LastChecked: time.Now(),
			Provider:    "dummy",
		},
		{
			Name:        "Thumbor",
			Description: "Image resizing service",
			Category:    "",
			Status:      ServiceStatusOk,
			LastChecked: time.Now(),
			Provider:    "dummy",
		},
	}
}
