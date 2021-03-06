// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010015 struct {
}

func (e *PG_E11010015) Error() string {
	return "NetID取引エラー 決済を中止して、取引ができない事を通知してください。"
}

func (e *PG_E11010015) Message() string {
	return "180日超えの取引のため、処理を行う事ができません。"
}

func (e *PG_E11010015) Code() string {
	return "E11010015"
}

func (e *PG_E11010015) CanRetry() bool {
	return false
}
