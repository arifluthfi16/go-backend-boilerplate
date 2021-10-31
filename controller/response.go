package controller

type ErrorResponse struct {
	Msg string
	Err error
}

type SuccessResponse struct {
	Status		bool	`json:"status"`
	Msg 		string 	`json:"msg"`
}
