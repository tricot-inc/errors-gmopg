// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01990006 struct{
}

func (e *PG_E01990006) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01990006) Message() string{
    return "過去6ヶ月間の購入回数に数字以外の文字が含まれています。"
}