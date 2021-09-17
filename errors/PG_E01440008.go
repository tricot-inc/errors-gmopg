// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01440008 struct{
}

func (e *PG_E01440008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01440008) Message() string{
    return "洗替・継続課金フラグの書式が正しくありません。"
}