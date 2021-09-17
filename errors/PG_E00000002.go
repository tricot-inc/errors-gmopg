// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E00000002 struct{
}

func (e *PG_E00000002) Error() string {
    return "接続方式エラー HTTPステータスコード406相当のエラーです。HTTPリクエストヘッダのAccept関連項目が正しくありません。"
}

func (e *PG_E00000002) Message() string {
    return "-"
}

func (e *PG_E00000002) CanRetry() bool {
    return false
}