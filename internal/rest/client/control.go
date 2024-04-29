package client

import (
	"context"

	"github.com/masnax/microcluster/internal/rest/types"
)

// ControlDaemon posts control data to the daemon.
func (c *Client) ControlDaemon(ctx context.Context, args types.Control) error {
	return c.QueryStruct(ctx, "POST", ControlEndpoint, nil, args, nil)
}
