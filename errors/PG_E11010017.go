// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010017 struct{
}

func (e *PG_E11010017) Error() string {
    return "リクルートかんたん支払い決済取引 エラー決済を中止して、取引ができない事を通知してください。"
}

func (e *PG_E11010017) Message() string {
    return "180日超えの取引のため、処理を行う事ができません。"
}

func (e *PG_E11010017) CanRetry() bool {
    return false
}