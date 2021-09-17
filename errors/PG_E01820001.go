// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01820001 struct{
}

func (e *PG_E01820001) Error() string {
    return "磁気ストライプ情報必須エラー 磁気ストライプ情報が指定されていません。"
}

func (e *PG_E01820001) Message() string {
    return "磁気ストライプ情報が不正です。"
}

func (e *PG_E01820001) CanRetry() bool {
    return false
}