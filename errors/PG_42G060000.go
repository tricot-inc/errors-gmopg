// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G060000 struct {
}

func (e *PG_42G060000) Error() string {
	return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G060000) Message() string {
	return "デビットカードで口座の残高が不足しています。"
}

func (e *PG_42G060000) Code() string {
	return "42G060000"
}

func (e *PG_42G060000) CanRetry() bool {
	return false
}
