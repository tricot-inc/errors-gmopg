// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01260002 struct {
}

func (e *PG_E01260002) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01260002) Message() string {
	return "指定された支払方法が存在しません。"
}

func (e *PG_E01260002) CanRetry() bool {
	return false
}
