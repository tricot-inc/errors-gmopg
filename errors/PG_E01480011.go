// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01480011 struct{
}

func (e *PG_E01480011) Error() string {
    return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01480011) Message() string {
    return "名義人の最大文字数を超えています。"
}

func (e *PG_E01480011) CanRetry() bool {
    return false
}