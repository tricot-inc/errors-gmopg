// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D80008 struct{
}

func (e *PG_E01D80008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D80008) Message() string{
    return "商品の販売時期情報の書式が正しくありません。"
}