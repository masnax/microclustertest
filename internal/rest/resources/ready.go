package resources

import (
	"fmt"
	"net/http"

	"github.com/canonical/lxd/lxd/response"

	"github.com/masnax/microclustertest/v2/internal/rest/access"
	"github.com/masnax/microclustertest/v2/internal/state"
	"github.com/masnax/microclustertest/v2/rest"
)

var readyCmd = rest.Endpoint{
	AllowedBeforeInit: true,
	Path:              "ready",

	Get: rest.EndpointAction{Handler: getWaitReady, AccessHandler: access.AllowAuthenticated},
}

func getWaitReady(state *state.State, r *http.Request) response.Response {
	if state.Context.Err() != nil {
		return response.Unavailable(fmt.Errorf("Daemon is shutting down"))
	}

	select {
	case <-state.ReadyCh:
	default:
		return response.Unavailable(fmt.Errorf("Daemon is not ready yet"))
	}

	return response.EmptySyncResponse
}
