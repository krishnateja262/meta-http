package metahttp

import (
	"context"
	"net/http"
)

type contextKey string

const (
	UserID           contextKey = "user-id"
	TenantID         contextKey = "tenant-id"
	RequestID        contextKey = "x-request-id"
	MerchantAPIKey   contextKey = "x-api-key"
	APIContextKey    contextKey = "apikey"
	AuthorizationKey contextKey = "Authorization"
)

var contextKeys = []contextKey{UserID, TenantID, RequestID, MerchantAPIKey, APIContextKey, AuthorizationKey}

func FetchHeadersFromContext(ctx context.Context) map[string]string {
	ctxHeaders := map[string]string{}
	for _, key := range contextKeys {
		val, ok := ctx.Value(key).(string)
		if ok {
			ctxHeaders[string(key)] = val
		}
	}

	return ctxHeaders
}

func FetchContextFromHeaders(ctx context.Context, r *http.Request) context.Context {
	for _, key := range contextKeys {
		val := r.Header.Get(string(key))
		if val != "" {
			ctx = context.WithValue(ctx, key, val)
		}
	}
	return ctx
}
