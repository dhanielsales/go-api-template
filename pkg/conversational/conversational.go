package conversational

import (
	"context"

	"github.com/google/uuid"
)

const (
	CID_CONTEXT_KEY = "cid"
	CID_HEADER_KEY  = "X-Conversational-ID"
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
