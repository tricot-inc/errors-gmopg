// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01870008 struct{
}

func (e *PG_E01870008) Error() string {
    return "トークンタイプの書式が正しくありません。"
}

func (e *PG_E01870008) Message() string {
    return "トークンタイプの書式が正しくありません。"
}

func (e *PG_E01870008) CanRetry() bool {
    return false
}