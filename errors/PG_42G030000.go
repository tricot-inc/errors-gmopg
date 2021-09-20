// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G030000 struct {
}

func (e *PG_42G030000) Error() string {
	return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G030000) Message() string {
	return "カード限度額を超えているために、決済を完了する事ができませんでした。"
}

func (e *PG_42G030000) Code() string {
	return "42G030000"
}

func (e *PG_42G030000) CanRetry() bool {
	return false
}
