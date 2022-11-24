package metahttp

import "context"

type contextKey string

const (
	UserID           contextKey = "user-id"
	TenantID         contextKey = "tenant-id"
	RequestID        contextKey = "request-id"
	MerchantAPIKey   contextKey = "x-api-key"
	APIContextKey    contextKey = "apikey"
	AuthorizationKey contextKey = "Authorization"
)

var contextKeys = []contextKey{UserID, TenantID, RequestID, MerchantAPIKey, APIContextKey, AuthorizationKey}

func fetchHeadersFromContext(ctx context.Context) map[string]string {
	ctxHeaders := map[string]string{}
	for _, key := range contextKeys {
		val, ok := ctx.Value(key).(string)
		if ok {
			ctxHeaders[string(key)] = val
		}
	}

	return ctxHeaders
}
