// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01270001 struct{
}

func (e *PG_E01270001) Error() string {
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01270001) Message() string {
    return "支払回数が指定されていません。"
}

func (e *PG_E01270001) CanRetry() bool {
    return false
}