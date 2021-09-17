// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01850008 struct {
}

func (e *PG_E01850008) Error() string {
	return "更新区分 更新区分の書式が正しくありません。"
}

func (e *PG_E01850008) Message() string {
	return "更新区分の書式が正しくありません。"
}

func (e *PG_E01850008) CanRetry() bool {
	return false
}
