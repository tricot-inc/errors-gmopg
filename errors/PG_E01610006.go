// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01610006 struct{
}

func (e *PG_E01610006) Error() string{
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01610006) Message() string{
    return "リトライ回数に数字以外が設定されています。"
}