// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01320013 struct{
}

func (e *PG_E01320013) Error() string{
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01320013) Message() string{
    return "加盟店自由項目1の値に利用できない文字が含まれています。"
}