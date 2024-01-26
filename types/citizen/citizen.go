package citizen

// CitizenI is the interface contract that citizen's keeper implements.
type CitizenI interface {
	GetCitizenID() string
	SetCitizenID(string) error
	GetCitizenName() string
	SetCitizenName(string) error
	GetEthAddress() string
	SetEthAddress(string) error
	GetAP() uint64
	SetAP(uint64) error
}
