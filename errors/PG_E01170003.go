// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01170003 struct{
}

func (e *PG_E01170003) Error() string{
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01170003) Message() string{
    return "カード番号が最大文字数を超えています。"
}