// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01490005 struct{
}

func (e *PG_E01490005) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01490005) Message() string {
    return "利用金額・税送料の合計値が有効な範囲を超えています。"
}

func (e *PG_E01490005) CanRetry() bool {
    return false
}