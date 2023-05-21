package structs

type Person struct {
	Name string
	Cpj  string
}

func (p Person) GetDocument() string {
	return p.Cpj
}
func (p Person) GetNickName() string {
	return p.Name
}
