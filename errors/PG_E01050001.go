// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01050001 struct{
}

func (e *PG_E01050001) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01050001) Message() string {
    return "処理区分が指定されていません。"
}

func (e *PG_E01050001) CanRetry() bool {
    return false
}