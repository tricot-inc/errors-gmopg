// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01350001 struct{
}

func (e *PG_E01350001) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01350001) Message() string {
    return "MDが指定されていません。"
}

func (e *PG_E01350001) CanRetry() bool {
    return false
}