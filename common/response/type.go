package response

type Responses interface {
	SetCode(int)
	GetCode() int
	SetTraceID(string)
	SetMsg(string)
	SetData(interface{})
	SetSuccess(bool)
	Clone() Responses
}
