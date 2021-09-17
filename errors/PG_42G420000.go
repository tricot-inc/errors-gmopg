// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G420000 struct {
}

func (e *PG_42G420000) Error() string {
	return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G420000) Message() string {
	return "暗証番号が誤っていた為に、決済を完了する事ができませんでした。"
}

func (e *PG_42G420000) CanRetry() bool {
	return false
}
