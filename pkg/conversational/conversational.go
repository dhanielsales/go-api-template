package conversational

import (
	"context"

	"github.com/dhanielsales/go-api-template/pkg/contextutils"
	"github.com/google/uuid"
)

const (
	CID_CONTEXT_KEY contextutils.ContextKey = "cid"                 // The key used for storing CID in the context.
	CID_HEADER_KEY  string                  = "X-Conversational-ID" // The header key used to send the CID in HTTP requests.
)

// NewCID generates a new Conversational ID (CID) using UUID.
func NewCID() string {
	return uuid.NewString()
}

// GetCIDFromContext retrieves the Conversational ID (CID) from the given context.
func GetCIDFromContext(ctx context.Context) string {
	if cid, ok := ctx.Value(CID_CONTEXT_KEY).(string); ok {
		return cid
	}

	return ""
}

// SetCIDToContext sets the Conversational ID (CID) in the given context.
func SetCIDToContext(ctx context.Context, cid string) context.Context {
	return context.WithValue(ctx, CID_CONTEXT_KEY, cid)
}
