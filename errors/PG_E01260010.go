// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01260010 struct{
}

func (e *PG_E01260010) Error() string {
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01260010) Message() string {
    return "指定されたカード番号または支払方法が正しくありません。"
}

func (e *PG_E01260010) CanRetry() bool {
    return false
}