package types

import (
	"github.com/masnax/microclustertest/rest/types"
)

// Server represents server status information.
type Server struct {
	Name    string         `json:"name"    yaml:"name"`
	Address types.AddrPort `json:"address" yaml:"address"`
	Ready   bool           `json:"ready"   yaml:"ready"`
}
