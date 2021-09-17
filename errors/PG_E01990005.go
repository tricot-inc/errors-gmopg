// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01990005 struct{
}

func (e *PG_E01990005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01990005) Message() string{
    return "過去6ヶ月間の購入回数が最大桁数を超えています。"
}