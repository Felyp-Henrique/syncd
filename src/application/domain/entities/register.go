package entities

type Register struct {
	Data map[string]any `json:"data"`
}

func NewRegister(data map[string]any) Register {
	return Register{
		Data: data,
	}
}
