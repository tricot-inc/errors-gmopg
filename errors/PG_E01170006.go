// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01170006 struct{
}

func (e *PG_E01170006) Error() string{
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01170006) Message() string{
    return "カード番号に数字以外の文字が含まれています。"
}