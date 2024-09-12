package conversational

import (
	"context"

	"github.com/dhanielsales/go-api-template/pkg/contextutils"
	"github.com/google/uuid"
)

const (
	CID_CONTEXT_KEY contextutils.ContextKey = "cid"
	CID_HEADER_KEY  string                  = "X-Conversational-ID"
)

func NewCID() string {
	return uuid.NewString()
}

func GetCIDFromContext(ctx context.Context) string {
	if cid, ok := ctx.Value(CID_CONTEXT_KEY).(string); ok {
		return cid
	}

	return ""
}

func SetCIDToContext(ctx context.Context, cid string) context.Context {
	return context.WithValue(ctx, CID_CONTEXT_KEY, cid)
}
