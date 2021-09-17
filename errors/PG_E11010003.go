// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010003 struct{
}

func (e *PG_E11010003) Error() string{
    return "取引エラー 決済を中止して、取引ができない事を通知してください。"
}

func (e *PG_E11010003) Message() string{
    return "この取引は指定処理区分処理を行う事ができません。"
}