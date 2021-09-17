// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010099 struct {
}

func (e *PG_E11010099) Error() string {
	return "取引エラー 使用されたカードの仕向先が判定できませんでした。"
}

func (e *PG_E11010099) Message() string {
	return "このカードはご利用になれません。"
}

func (e *PG_E11010099) CanRetry() bool {
	return false
}
