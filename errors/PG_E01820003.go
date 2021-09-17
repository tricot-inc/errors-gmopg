// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01820003 struct {
}

func (e *PG_E01820003) Error() string {
	return "磁気ストライプ情報文字数エラー 磁気ストライプ情報の文字数が正しくありません。"
}

func (e *PG_E01820003) Message() string {
	return "磁気ストライプ情報が不正です。"
}

func (e *PG_E01820003) CanRetry() bool {
	return false
}
