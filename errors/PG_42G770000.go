// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G770000 struct{
}

func (e *PG_42G770000) Error() string {
    return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。詳細はお問い合わせください。"
}

func (e *PG_42G770000) Message() string {
    return "業務区分に誤りがあるために、決済を完了する事ができませんでした。"
}

func (e *PG_42G770000) CanRetry() bool {
    return false
}