// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B60005 struct{
}

func (e *PG_E01B60005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B60005) Message() string{
    return "請求先住所の郵便番号が最大桁数を超えています。"
}