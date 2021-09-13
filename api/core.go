package api

import "context"

// CoreCloudRuntimes is the interface for cloud-runtimes-api.
type CoreCloudRuntimes interface {
	InvocationRuntimes
	BindingRuntimes
	PubSubRuntimes
	SecretsRuntimes
	StateRuntimes
}

// -------------------------- invoke

type InvocationRuntimes interface {

	// InvokeMethod invokes service without raw data
	InvokeMethod(ctx context.Context, appID, methodName, verb string) (out []byte, err error)

	// InvokeMethodWithContent invokes service with content
	InvokeMethodWithContent(ctx context.Context, appID, methodName, verb string, content *DataContent) (out []byte, err error)

	// InvokeMethodWithCustomContent invokes app with custom content (struct + content type).
	InvokeMethodWithCustomContent(ctx context.Context, appID, methodName, verb string, contentType string, content interface{}) (out []byte, err error)
}

// DataContent the service invocation content
type DataContent struct {
	// Data is the input data
	Data []byte
	// ContentType is the type of the data content
	ContentType string
}

// -------------------------- binding

type BindingRuntimes interface {

	// InvokeBinding invokes specific operation on the configured binding.
	// This method covers input, output, and bi-directional bindings.
	InvokeBinding(ctx context.Context, in *InvokeBindingRequest) (out *BindingEvent, err error)

	// InvokeOutputBinding invokes configured binding with data.InvokeOutputBinding
	// This method differs from InvokeBinding in that it doesn't expect any content being returned from the invoked method.
	InvokeOutputBinding(ctx context.Context, in *InvokeBindingRequest) error
}

// InvokeBindingRequest represents binding invocation request
type InvokeBindingRequest struct {
	// Name is name of binding to invoke.
	Name string
	// Operation is the name of the operation type for the binding to invoke
	Operation string
	// Data is the input bindings sent
	Data []byte
	// Metadata is the input binding metadata
	Metadata map[string]string
}

// BindingEvent represents the binding event handler input
type BindingEvent struct {
	// Data is the input bindings sent
	Data []byte
	// Metadata is the input binding metadata
	Metadata map[string]string
}

// -------------------------- pub/sub

type PubSubRuntimes interface {

	// PublishEvent publishes data onto topic in specific pubsub component.
	PublishEvent(ctx context.Context, pubsubName, topicName string, data []byte) error

	// PublishEventFromCustomContent serializes an struct and publishes its contents as data (JSON) onto topic in specific pubsub component.
	PublishEventFromCustomContent(ctx context.Context, pubsubName, topicName string, data interface{}) error
}

// -------------------------- secrets

type SecretsRuntimes interface {

	// GetSecret retrieves preconfigured secret from specified store using key.
	GetSecret(ctx context.Context, storeName, key string, meta map[string]string) (data map[string]string, err error)

	// GetBulkSecret retrieves all preconfigured secrets for this application.
	GetBulkSecret(ctx context.Context, storeName string, meta map[string]string) (data map[string]map[string]string, err error)
}

// -------------------------- state

type StateRuntimes interface {

	// SaveState saves the raw data into store using default state options.
	SaveState(ctx context.Context, storeName, key string, data []byte, so ...StateOption) error

	// SaveBulkState saves multiple state item to store with specified options.
	SaveBulkState(ctx context.Context, storeName string, items ...*SetStateItem) error

	// GetState retrieves state from specific store using default consistency option.
	GetState(ctx context.Context, storeName, key string) (item *StateItem, err error)

	// GetStateWithConsistency retrieves state from specific store using provided state consistency.
	GetStateWithConsistency(ctx context.Context, storeName, key string, meta map[string]string, sc StateConsistency) (item *StateItem, err error)

	// GetBulkState retrieves state for multiple keys from specific store.
	GetBulkState(ctx context.Context, storeName string, keys []string, meta map[string]string, parallelism int32) ([]*BulkStateItem, error)

	// DeleteState deletes content from store using default state options.
	DeleteState(ctx context.Context, storeName, key string) error

	// DeleteStateWithETag deletes content from store using provided state options and etag.
	DeleteStateWithETag(ctx context.Context, storeName, key string, etag *ETag, meta map[string]string, opts *StateOptions) error

	// ExecuteStateTransaction provides way to execute multiple operations on a specified store.
	ExecuteStateTransaction(ctx context.Context, storeName string, meta map[string]string, ops []*StateOperation) error

	// DeleteBulkState deletes content for multiple keys from store.
	DeleteBulkState(ctx context.Context, storeName string, keys []string) error

	// DeleteBulkState deletes content for multiple keys from store.
	DeleteBulkStateItems(ctx context.Context, storeName string, items []*DeleteStateItem) error
}

type (
	// StateConsistency is the consistency enum type.
	StateConsistency int
	// StateConcurrency is the concurrency enum type.
	StateConcurrency int
	// OperationType is the operation enum type.
	OperationType int
)

// StateOperation is a collection of StateItems with a store name.
type StateOperation struct {
	Type OperationType
	Item *SetStateItem
}

// StateItem represents a single state item.
type StateItem struct {
	Key      string
	Value    []byte
	Etag     string
	Metadata map[string]string
}

// BulkStateItem represents a single state item.
type BulkStateItem struct {
	Key      string
	Value    []byte
	Etag     string
	Metadata map[string]string
	Error    string
}

// SetStateItem represents a single state to be persisted.
type SetStateItem struct {
	Key      string
	Value    []byte
	Etag     *ETag
	Metadata map[string]string
	Options  *StateOptions
}

// DeleteStateItem represents a single state to be deleted.
type DeleteStateItem SetStateItem

// ETag represents an versioned record information
type ETag struct {
	Value string
}

// StateOptions represents the state store persistence policy.
type StateOptions struct {
	Concurrency StateConcurrency
	Consistency StateConsistency
}

// StateOption StateOptions's function type
type StateOption func(*StateOptions)
