// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G830000 struct{
}

func (e *PG_42G830000) Error() string{
    return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。"
}

func (e *PG_42G830000) Message() string{
    return "有効期限に誤りがあるために、決済を完了する事ができませんでした。"
}