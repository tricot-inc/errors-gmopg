// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B20005 struct{
}

func (e *PG_E01B20005) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B20005) Message() string {
    return "請求先住所の国番号が3桁ではありません。"
}

func (e *PG_E01B20005) CanRetry() bool {
    return false
}