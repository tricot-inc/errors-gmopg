// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01040010 struct{
}

func (e *PG_E01040010) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01040010) Message() string{
    return "すでにオーダーIDが存在しています。"
}