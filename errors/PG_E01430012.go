// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01430012 struct {
}

func (e *PG_E01430012) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01430012) Message() string {
	return "会員名の値が最大バイト数を超えています。"
}

func (e *PG_E01430012) Code() string {
	return "E01430012"
}

func (e *PG_E01430012) CanRetry() bool {
	return false
}
