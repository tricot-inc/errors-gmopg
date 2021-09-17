// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01060021 struct{
}

func (e *PG_E01060021) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01060021) Message() string{
    return "取引の利用金額と指定した利用金額が一致していません。"
}