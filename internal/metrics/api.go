package metrics

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterHandlers registers the handlers that perform healthchecks.
func RegisterHandlers(r *routing.Router, version string) {
	r.To("GET,HEAD", "/metrics", routing.HTTPHandler(promhttp.Handler()))
}
