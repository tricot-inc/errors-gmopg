// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01060011 struct{
}

func (e *PG_E01060011) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01060011) Message() string {
    return "利用金額が有効な範囲を超えています。"
}

func (e *PG_E01060011) CanRetry() bool {
    return false
}