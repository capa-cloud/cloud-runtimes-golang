package api

import "context"

// EnhancedCloudRuntimes is the interface for cloud-runtimes-api.
type EnhancedCloudRuntimes interface {
	TraceRuntimes
	AuthRuntimes
}

// -------------------------- trace

type TraceRuntimes interface {

	// WithTraceID adds existing trace ID to the outgoing context.
	WithTraceID(ctx context.Context, id string) context.Context
}

// -------------------------- auth

type AuthRuntimes interface {

	// WithAuthToken sets auth API token on the instantiated client.
	WithAuthToken(token string)
}
