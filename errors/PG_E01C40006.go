// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C40006 struct{
}

func (e *PG_E01C40006) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C40006) Message() string{
    return "職場電話番号に数字以外の文字が含まれています。"
}