// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G300000 struct {
}

func (e *PG_42G300000) Error() string {
	return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G300000) Message() string {
	return "このカードでは取引をする事ができません。"
}

func (e *PG_42G300000) Code() string {
	return "42G300000"
}

func (e *PG_42G300000) CanRetry() bool {
	return false
}
