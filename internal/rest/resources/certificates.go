package resources

import (
	"context"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/canonical/lxd/lxd/response"

	"github.com/masnax/microclustertest/v0/client"
	"github.com/masnax/microclustertest/v0/internal/rest/access"
	"github.com/masnax/microclustertest/v0/internal/state"
	"github.com/masnax/microclustertest/v0/rest"
	"github.com/masnax/microclustertest/v0/rest/types"
)

var clusterCertificatesCmd = rest.Endpoint{
	AllowedBeforeInit: true,
	Path:              "cluster/certificates",

	Put: rest.EndpointAction{Handler: clusterCertificatesPut, AccessHandler: access.AllowAuthenticated},
}

func clusterCertificatesPut(s *state.State, r *http.Request) response.Response {
	req := types.ClusterCertificatePut{}

	// Parse the request.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return response.BadRequest(err)
	}

	// Forward the request to all other nodes if we are the first.
	if !client.IsNotification(r) && s.Database.IsOpen() {
		cluster, err := s.Cluster(true)
		if err != nil {
			return response.SmartError(err)
		}

		err = cluster.Query(s.Context, true, func(ctx context.Context, c *client.Client) error {
			return c.UpdateClusterCertificate(ctx, req)
		})
		if err != nil {
			return response.SmartError(fmt.Errorf("Failed to update cluster certificate on peers: %w", err))
		}
	}

	certBlock, _ := pem.Decode([]byte(req.PublicKey))
	if certBlock == nil {
		return response.BadRequest(fmt.Errorf("Certificate must be base64 encoded PEM certificate"))
	}

	keyBlock, _ := pem.Decode([]byte(req.PrivateKey))
	if keyBlock == nil {
		return response.BadRequest(fmt.Errorf("Private key must be base64 encoded PEM key"))
	}

	// If a CA was specified, validate that as well.
	if req.CA != "" {
		caBlock, _ := pem.Decode([]byte(req.CA))
		if caBlock == nil {
			return response.BadRequest(fmt.Errorf("CA must be base64 encoded PEM key"))
		}

		err = os.WriteFile(filepath.Join(s.OS.StateDir, "cluster.ca"), []byte(req.CA), 0650)
		if err != nil {
			return response.SmartError(err)
		}
	}

	// Write the keypair to the state directory.
	err = os.WriteFile(filepath.Join(s.OS.StateDir, "cluster.crt"), []byte(req.PublicKey), 0650)
	if err != nil {
		return response.SmartError(err)
	}

	err = os.WriteFile(filepath.Join(s.OS.StateDir, "cluster.key"), []byte(req.PrivateKey), 0650)
	if err != nil {
		return response.SmartError(err)
	}

	// Load the new cluster cert from the state directory on this node.
	err = state.ReloadClusterCert()
	if err != nil {
		return response.SmartError(err)
	}

	return response.EmptySyncResponse
}
