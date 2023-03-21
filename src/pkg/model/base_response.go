package model

type BaseResponseModel[T_DATA any] struct {
	Data           T_DATA `json:"data,omitempty"`
	Message        string `json:"message"`
	Status         string `json:"status"`
	ErrorMessage   string `json:"-"`
	StatusCode     string `json:"-"`
	HttpStatusCode int    `json:"-"`
	PaginationResponseModel
}

func (m *BaseResponseModel[T_DATA]) Set(status, message string, httpStatus int, data T_DATA) {
	m.Status = status
	m.Message = message
	m.HttpStatusCode = httpStatus
}

func (m *BaseResponseModel[T_DATA]) SetData(data T_DATA) *BaseResponseModel[T_DATA] {
	m.Data = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetMessage(data string) *BaseResponseModel[T_DATA] {
	m.Message = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetErrorMessage(data string) *BaseResponseModel[T_DATA] {
	m.ErrorMessage = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetSuccess(data string) *BaseResponseModel[T_DATA] {
	m.Status = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetStatusCode(data string) *BaseResponseModel[T_DATA] {
	m.StatusCode = data
	return m
}

func (m *BaseResponseModel[T_DATA]) SetHttpStatusCode(data int) *BaseResponseModel[T_DATA] {
	m.HttpStatusCode = data
	return m
}
