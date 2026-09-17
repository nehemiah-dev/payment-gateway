package api

type HealthStatus string

const (
	HealthStatusHealthy   HealthStatus = "healthy"
	HealthStatusUnhealthy HealthStatus = "unhealthy"
)

type HealthResponse struct {
	Status *HealthStatus `json:"status,omitempty"`
}

func (hs HealthStatus) String() string {
	return string(hs)
}

func (hs HealthStatus) Ptr() *HealthStatus {
	v := hs
	return &v
}
