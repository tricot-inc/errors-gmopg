// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C80005 struct{
}

func (e *PG_E01C80005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C80005) Message() string{
    return "配送先住所の区域部分の2行目が最大桁数を超えています。"
}