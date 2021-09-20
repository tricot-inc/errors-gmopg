// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01160010 struct {
}

func (e *PG_E01160010) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01160010) Message() string {
	return "ボーナス分割回数に“2”以外を指定しています。"
}

func (e *PG_E01160010) Code() string {
	return "E01160010"
}

func (e *PG_E01160010) CanRetry() bool {
	return false
}
