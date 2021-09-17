// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01180008 struct {
}

func (e *PG_E01180008) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01180008) Message() string {
	return "有効期限の書式が正しくありません。"
}

func (e *PG_E01180008) CanRetry() bool {
	return false
}
