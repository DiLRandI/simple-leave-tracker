package model

func GetModels() []any {
	return []any{
		&User{},
		&Login{},
	}
}
