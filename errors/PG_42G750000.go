// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G750000 struct{
}

func (e *PG_42G750000) Error() string{
    return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。詳細はカード会社にお問い合わせください。"
}

func (e *PG_42G750000) Message() string{
    return "分割金額に誤りがあるために、決済を完了する事ができませんでした。"
}