package rest

import (
	"net/http"

	"github.com/canonical/lxd/lxd/response"

	"github.com/masnax/microclustertest/v1/state"
)

// EndpointAlias represents an alias URL of and Endpoint in our API.
type EndpointAlias struct {
	Name string // Name for this alias.
	Path string // Path pattern for this alias.
}

// EndpointAction represents an action on an API endpoint.
type EndpointAction struct {
	Handler        func(state *state.State, r *http.Request) response.Response
	AccessHandler  func(state *state.State, r *http.Request) response.Response
	AllowUntrusted bool
	ProxyTarget    bool // Allow forwarding of the request to a target if ?target=name is specified.
}

// Endpoint represents a URL in our API.
type Endpoint struct {
	Name    string          // Name for this endpoint.
	Path    string          // Path pattern for this endpoint.
	Aliases []EndpointAlias // Any aliases for this endpoint.
	Get     EndpointAction
	Put     EndpointAction
	Post    EndpointAction
	Delete  EndpointAction
	Patch   EndpointAction

	AllowedDuringShutdown bool // Whether we should return Unavailable Error (503) if daemon is shutting down.
	AllowedBeforeInit     bool // Whether we should return Unavailabel Error (503) if the daemon has not been initialized (is not yet part of a cluster).
}
