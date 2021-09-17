// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42G810000 struct{
}

func (e *PG_42G810000) Error() string{
    return "カード会社オーソリエラー 指定されたカードのオーソリが失敗した事を通知してください。詳細はお問い合わせください。"
}

func (e *PG_42G810000) Message() string{
    return "取消取扱区分に誤りがあるために、決済を完了する事ができませんでした。"
}