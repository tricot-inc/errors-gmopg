// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01810008 struct {
}

func (e *PG_E01810008) Error() string {
	return "磁気ストライプ区分書式エラー 磁気ストライプ区分に不正値が設定されています。"
}

func (e *PG_E01810008) Message() string {
	return "磁気ストライプ区分が不正です。"
}

func (e *PG_E01810008) CanRetry() bool {
	return false
}
