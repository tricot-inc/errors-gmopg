// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B20002 struct{
}

func (e *PG_E01B20002) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B20002) Message() string {
    return "請求先住所の国番号が存在しません。"
}

func (e *PG_E01B20002) CanRetry() bool {
    return false
}