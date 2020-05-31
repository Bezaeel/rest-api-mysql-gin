package ApiResponse

type Response struct {
	IsSuccess bool
	Message   string
	Data      interface{}
}
