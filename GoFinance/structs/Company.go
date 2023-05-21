package structs

type Company struct {
	FantasyName string
	Cnpj        string
}

func (c Company) GetDocument() string {
	return c.Cnpj
}
func (c Company) GetNickName() string {
	return c.FantasyName
}
