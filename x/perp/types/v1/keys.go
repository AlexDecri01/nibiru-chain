package types

var (
	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	MemStoreKey = "mem_" + ModuleName

	// RouterKey is the message route for perp.
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName
)
