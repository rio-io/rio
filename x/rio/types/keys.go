package types

const (
	// ModuleName defines the module name
	ModuleName = "rio"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rio"

	// Keep track of the index of certs
	CertKey = "Cert-value-"

	// TODO: It really need it? not unique key combined with creator and sequential id
	CertCountKey = "Cert-count-"

	ResumeKey      = "Resume-value-"
	ResumeCountKey = "Resume-count-"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
