// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C70005 struct {
}

func (e *PG_E01C70005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C70005) Message() string {
	return "配送先住所の区域部分の1行目が最大桁数を超えています。"
}

func (e *PG_E01C70005) Code() string {
	return "E01C70005"
}

func (e *PG_E01C70005) CanRetry() bool {
	return false
}
