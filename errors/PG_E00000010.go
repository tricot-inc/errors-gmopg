// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E00000010 struct{
}

func (e *PG_E00000010) Error() string {
    return "接続方式エラー HTTPコンテンツが空、またはフォーマットが正しくありません。"
}

func (e *PG_E00000010) Message() string {
    return "-"
}

func (e *PG_E00000010) CanRetry() bool {
    return false
}