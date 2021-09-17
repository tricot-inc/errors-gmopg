// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01810001 struct{
}

func (e *PG_E01810001) Error() string {
    return "磁気ストライプ区分必須エラー 磁気ストライプ区分が指定されていません"
}

func (e *PG_E01810001) Message() string {
    return "磁気ストライプ区分が不正です。"
}

func (e *PG_E01810001) CanRetry() bool {
    return false
}