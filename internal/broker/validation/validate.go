package validation

type Validator struct {
}

type Error struct {
	Name string
	Msg  string
}

type Errors struct {
	_ []Error
}

func NewValidator() *Validator {
	return &Validator{}
}
