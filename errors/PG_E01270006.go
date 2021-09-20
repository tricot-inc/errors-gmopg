// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01270006 struct {
}

func (e *PG_E01270006) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01270006) Message() string {
	return "支払回数の数字以外の文字が含まれています。"
}

func (e *PG_E01270006) Code() string {
	return "E01270006"
}

func (e *PG_E01270006) CanRetry() bool {
	return false
}
