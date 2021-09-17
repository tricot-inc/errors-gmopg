// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A50005 struct{
}

func (e *PG_E01A50005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A50005) Message() string{
    return "過去24時間の取引回数が最大桁数を超えています。"
}