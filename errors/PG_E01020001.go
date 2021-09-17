// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01020001 struct{
}

func (e *PG_E01020001) Error() string {
    return "入力パラメータエラー 定を確認してください。"
}

func (e *PG_E01020001) Message() string {
    return "ショップパスワードが指定されていません。"
}

func (e *PG_E01020001) CanRetry() bool {
    return false
}