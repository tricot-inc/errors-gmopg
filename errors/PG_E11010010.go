// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010010 struct{
}

func (e *PG_E11010010) Error() string{
    return "カード取引エラー 決済を中止して、取引ができない事を通知してください。"
}

func (e *PG_E11010010) Message() string{
    return "180日超えの取引のため、処理を行う事ができません。"
}