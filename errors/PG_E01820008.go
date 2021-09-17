// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01820008 struct {
}

func (e *PG_E01820008) Error() string {
	return "磁気ストライプ情報書式エラー 磁気ストライプ情報の書式が正しくありません。"
}

func (e *PG_E01820008) Message() string {
	return "磁気ストライプ情報が不正です。"
}

func (e *PG_E01820008) Code() string {
	return "E01820008"
}

func (e *PG_E01820008) CanRetry() bool {
	return false
}
