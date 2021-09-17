// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E80001 struct{
}

func (e *PG_E01E80001) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E80001) Message() string{
    return "携帯電話番号の指定が正しくありません。携帯電話の国コード/携帯電話番号のいずれかの省略はできません。"
}