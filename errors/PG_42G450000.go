// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G450000 struct {
}

func (e *PG_42G450000) Error() string {
	return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G450000) Message() string {
	return "セキュリティーコードが入力されていない為に、決済を完了する事ができませんでした。"
}

func (e *PG_42G450000) Code() string {
	return "42G450000"
}

func (e *PG_42G450000) CanRetry() bool {
	return false
}
