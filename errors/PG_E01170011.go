// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01170011 struct {
}

func (e *PG_E01170011) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01170011) Message() string {
	return "カード番号が10桁~16桁の範囲ではありません。"
}

func (e *PG_E01170011) Code() string {
	return "E01170011"
}

func (e *PG_E01170011) CanRetry() bool {
	return false
}
