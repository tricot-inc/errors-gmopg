// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A20008 struct{
}

func (e *PG_E01A20008) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A20008) Message() string {
    return "配送先住所の初回使用日の書式が正しくありません。"
}

func (e *PG_E01A20008) CanRetry() bool {
    return false
}