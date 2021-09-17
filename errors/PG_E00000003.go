// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E00000003 struct {
}

func (e *PG_E00000003) Error() string {
	return "接続方式エラー HTTPステータスコード415相当のエラーです。HTTPリクエストヘッダのContent-type項目が正しくありません。"
}

func (e *PG_E00000003) Message() string {
	return "-"
}

func (e *PG_E00000003) Code() string {
	return "E00000003"
}

func (e *PG_E00000003) CanRetry() bool {
	return false
}
