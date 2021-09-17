// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D30008 struct{
}

func (e *PG_E01D30008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D30008) Message() string{
    return "商品納品時間枠の書式が正しくありません。"
}