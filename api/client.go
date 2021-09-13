package api

import "context"

// CloudRuntimesClient is the interface for cloud-runtimes client implementation.
type CloudRuntimesClient interface {
	CoreCloudRuntimes
	EnhancedCloudRuntimes

	// Shutdown Gracefully shutdown the cloud-runtimes runtime.
	Shutdown(ctx context.Context) error

	// Close cleans up all resources created by the client.
	Close()
}
