// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D40005 struct{
}

func (e *PG_E01D40005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D40005) Message() string{
    return "プリペイドカードまたはギフトカードの総購入金額が最大桁数を超えています。"
}