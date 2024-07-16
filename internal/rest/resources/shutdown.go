package resources

import (
	"fmt"
	"net/http"

	"github.com/canonical/lxd/lxd/response"

	"github.com/masnax/microclustertest/v5/internal/rest/access"
	"github.com/masnax/microclustertest/v5/internal/state"
	"github.com/masnax/microclustertest/v5/rest"
)

var shutdownCmd = rest.Endpoint{
	AllowedBeforeInit: true,
	Path:              "shutdown",

	Post: rest.EndpointAction{Handler: shutdownPost, AccessHandler: access.AllowAuthenticated},
}

func shutdownPost(state *state.State, r *http.Request) response.Response {
	if state.Context.Err() != nil {
		return response.SmartError(fmt.Errorf("Shutdown already in progress"))
	}

	return response.ManualResponse(func(w http.ResponseWriter) error {
		<-state.ReadyCh // Wait for daemon to start.

		// Run shutdown sequence synchronously.
		exit, stopErr := state.Stop()
		err := response.SmartError(stopErr).Render(w)
		if err != nil {
			return err
		}

		// Send the response before the daemon process ends.
		f, ok := w.(http.Flusher)
		if ok {
			f.Flush()
		} else {
			return fmt.Errorf("ResponseWriter is not type http.Flusher")
		}

		// Send result of d.Stop() to cmdDaemon so that process stops with correct exit code from Stop().
		go func() {
			<-r.Context().Done() // Wait until request is finished.
			exit()
		}()

		return nil
	})
}
