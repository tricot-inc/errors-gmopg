// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G070000 struct {
}

func (e *PG_42G070000) Error() string {
	return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G070000) Message() string {
	return "カード限度額を超えているために、決済を完了する事ができませんでした。"
}

func (e *PG_42G070000) Code() string {
	return "42G070000"
}

func (e *PG_42G070000) CanRetry() bool {
	return false
}
