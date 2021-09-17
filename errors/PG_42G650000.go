// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G650000 struct{
}

func (e *PG_42G650000) Error() string {
    return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。カード番号の再入力依頼をしてください。"
}

func (e *PG_42G650000) Message() string {
    return "カード番号に誤りがあるために、決済を完了する事ができませんでした。"
}

func (e *PG_42G650000) CanRetry() bool {
    return false
}