// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01230009 struct{
}

func (e *PG_E01230009) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01230009) Message() string {
    return "カード登録連番が最大登録可能数を超えています。"
}

func (e *PG_E01230009) CanRetry() bool {
    return false
}