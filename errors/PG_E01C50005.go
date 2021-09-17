// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C50005 struct {
}

func (e *PG_E01C50005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C50005) Message() string {
	return "配送先住所の都市が最大桁数を超えています。"
}

func (e *PG_E01C50005) Code() string {
	return "E01C50005"
}

func (e *PG_E01C50005) CanRetry() bool {
	return false
}
