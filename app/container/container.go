// package container use dependency injection to create concrete type and wire the whole application together
package container

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file (if you have configuration file).
// Client app use those to retrieve use case from the container
const (
	SEARCH_ORDER_USECASE string = "GetOrderUseCase"
	MODIFY_ORDER_USECASE string = "ModifyOrderUseCase"

	DATABASE         string = "database"
	EVENT_BUS        string = "eventBus"
	MESSAGING_SERVER string = "messagingServer"

)

type Container interface {
	// BuildUseCase creates concrete types for use case and it's included types.
	// For each call, it will create a new instance, which means it is not a singleton
	BuildUseCase(code string) (interface{}, error)

	// This should only be used by container and it's sub-package
	// Get instance by code from container.
	Get(code string) (interface{}, bool)

	// This should only be used by container and it's sub-package
	// Put value into container with code as the key.
	Put(code string, value interface{})

}

