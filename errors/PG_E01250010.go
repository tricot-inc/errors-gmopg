// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01250010 struct {
}

func (e *PG_E01250010) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01250010) Message() string {
	return "カードパスワードが一致しません。"
}

func (e *PG_E01250010) Code() string {
	return "E01250010"
}

func (e *PG_E01250010) CanRetry() bool {
	return false
}
