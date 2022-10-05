package errorx

type ServiceError struct {
	Status int    `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (se *ServiceError) Error() string {
	return se.Msg
}

var (
	InternalServerError       = &ServiceError{500, 50000, "internal server error"}
	CreateHashIdFailedError   = &ServiceError{500, 50001, "internal server error"}
	InsertDataBaseFailedError = &ServiceError{500, 50002, "internal server error"}
	FetchDatabaseFailedError  = &ServiceError{500, 50003, "internal server error"}

	DataNotFoundError = &ServiceError{404, 40004, "data not found"}

	BadRequestError       = &ServiceError{400, 40001, "bad request"}
	InvalidParameterError = &ServiceError{400, 40002, "invalid parameter"}
	UrlNotFoundError      = &ServiceError{404, 40400, "url not found"}
)
