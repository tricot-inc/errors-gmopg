// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01480008 struct{
}

func (e *PG_E01480008) Error() string{
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01480008) Message() string{
    return "名義人の書式が正しくありません。"
}