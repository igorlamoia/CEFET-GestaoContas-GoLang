package structs

type OwnerType string

const (
	PF OwnerType = "PF"
	PJ OwnerType = "PJ"
)

type Owner interface {
	GetDocument() string
	GetNickName() string
	GetOwnerType() OwnerType
}

func (c Company) GetOwnerType() OwnerType {
	return PJ
}
func (p Person) GetOwnerType() OwnerType {
	return PF
}