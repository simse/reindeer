package status

import "time"

type Service struct {
	ID          int    `storm:"id,increment"`
	Name        string `storm:"index"`
	Description string
	Category    string
	URL         string
	Status      ServiceStatus
	LastChecked time.Time `storm:"index"`
	Provider    string
}

type ServiceStatus string

const (
	ServiceStatusOk    ServiceStatus = "OK"
	ServiceStatusWarn  ServiceStatus = "ISSUE"
	ServiceStatusError ServiceStatus = "ERROR"
)

type Summary struct {
	LastChecked time.Time
	Status      ServiceStatus
}

func GetSummary() Summary {
	now := time.Now()
	lastChecked := now.Add(time.Duration(-2) * time.Minute)

	return Summary{
		LastChecked: lastChecked,
		Status:      ServiceStatusOk,
	}
}
