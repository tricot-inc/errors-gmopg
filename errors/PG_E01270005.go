// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01270005 struct {
}

func (e *PG_E01270005) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01270005) Message() string {
	return "支払回数が最大桁数を超えています。"
}

func (e *PG_E01270005) CanRetry() bool {
	return false
}
