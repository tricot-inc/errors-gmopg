// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E91099996 struct{
}

func (e *PG_E91099996) Error() string{
    return "システムエラー(通信)"
}

func (e *PG_E91099996) Message() string{
    return "システムの内部エラーです。発生時刻や呼び出しパラメータをご確認のうえ、お問い合わせください。"
}