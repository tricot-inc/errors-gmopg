// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E20006 struct{
}

func (e *PG_E01E20006) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E20006) Message() string {
    return "継続課金の課金最小間隔日数に数字以外の文字が含まれています。"
}

func (e *PG_E01E20006) CanRetry() bool {
    return false
}