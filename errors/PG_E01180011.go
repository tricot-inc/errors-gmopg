// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01180011 struct{
}

func (e *PG_E01180011) Error() string{
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01180011) Message() string{
    return "有効期限の書式が正しくありません。"
}