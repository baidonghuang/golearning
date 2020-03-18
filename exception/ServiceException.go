package exception

/**
自定义异常
created by hbd
2020-03-13
*/
type ServiceException struct {
	ErrorMsg  string
	ErrorCode int
}

//实现error的Error函数
func (e *ServiceException) Error() string {
	return e.ErrorMsg
}

func (e *ServiceException) getErrorCode() int {
	return e.ErrorCode
}
