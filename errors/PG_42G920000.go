// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G920000 struct {
}

func (e *PG_42G920000) Error() string {
	return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G920000) Message() string {
	return "このカードでは取引をする事ができません。"
}

func (e *PG_42G920000) Code() string {
	return "42G920000"
}

func (e *PG_42G920000) CanRetry() bool {
	return false
}
