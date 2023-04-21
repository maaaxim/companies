package api

type errorsResponse struct {
	Success bool     `json:"success" example:"false"` // always false
	Error   []string `json:"error"`                   // error descriptions array
}

func newErrorsResponse(errs []error) errorsResponse {
	errorAsStrings := make([]string, 0, len(errs))
	for _, err := range errs {
		errorAsStrings = append(errorAsStrings, err.Error())
	}

	return errorsResponse{
		Success: false,
		Error:   errorAsStrings,
	}
}

type successResponse struct {
	Success bool        `json:"success" example:"true"` // always true
	Data    interface{} `json:"data"`                   // object or empty object
}

func newSuccessResponse(jsonableResponse any) successResponse {
	return successResponse{
		Success: true,
		Data:    jsonableResponse,
	}
}
