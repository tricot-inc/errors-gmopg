// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01260001 struct {
}

func (e *PG_E01260001) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01260001) Message() string {
	return "支払方法が指定されていません。"
}

func (e *PG_E01260001) Code() string {
	return "E01260001"
}

func (e *PG_E01260001) CanRetry() bool {
	return false
}
