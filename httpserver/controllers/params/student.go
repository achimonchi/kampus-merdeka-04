package params

type StudentCreateRequest struct {
	Name  string `validate:"required"`
	Age   int    `validate:"required"`
	Grade int    `validate:"required"`
}
