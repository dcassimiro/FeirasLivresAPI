package feira

type searchFeira struct {
	Distrito string `json:"distrito" validate:"required"`
}
