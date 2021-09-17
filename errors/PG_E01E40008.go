// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E40008 struct{
}

func (e *PG_E01E40008) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E40008) Message() string {
    return "認証状態の書式が正しくありません。"
}

func (e *PG_E01E40008) CanRetry() bool {
    return false
}