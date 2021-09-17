// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E00000001 struct {
}

func (e *PG_E00000001) Error() string {
	return "接続方式エラー HTTPステータスコード405相当のエラーです。HTTPメソッドがPOSTではありません。"
}

func (e *PG_E00000001) Message() string {
	return "-"
}

func (e *PG_E00000001) Code() string {
	return "E00000001"
}

func (e *PG_E00000001) CanRetry() bool {
	return false
}
