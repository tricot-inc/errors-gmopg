// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01070006 struct{
}

func (e *PG_E01070006) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01070006) Message() string{
    return "税送料に数字以外の文字が含まれています。"
}