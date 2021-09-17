// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01460008 struct{
}

func (e *PG_E01460008) Error() string{
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01460008) Message() string{
    return "セキュリティコードの書式が正しくありません。"
}