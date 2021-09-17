// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C60005 struct{
}

func (e *PG_E01C60005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C60005) Message() string{
    return "配送先住所の国番号が3桁ではありません。"
}