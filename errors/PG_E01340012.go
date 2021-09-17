// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01340012 struct{
}

func (e *PG_E01340012) Error() string{
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01340012) Message() string{
    return "加盟店自由項目3の値が最大バイト数を超えています。"
}