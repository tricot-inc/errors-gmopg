// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D10005 struct{
}

func (e *PG_E01D10005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D10005) Message() string{
    return "配送先住所の州または都道府県番号が最大桁数を超えています。"
}