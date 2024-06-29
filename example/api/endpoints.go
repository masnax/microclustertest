// Package api provides the REST API endpoints.
package api

import (
	"github.com/masnax/microclustertest/v2/rest"
)

// Endpoints is a global list of all API endpoints on the /1.0 endpoint of
// microcluster, as supplied by this example project.
var Endpoints = []rest.Endpoint{
	extendedCmd,
}
