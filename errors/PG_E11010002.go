// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010002 struct {
}

func (e *PG_E11010002) Error() string {
	return "取引エラー 決済を中止して、取引ができない事を通知してください。"
}

func (e *PG_E11010002) Message() string {
	return "この取引は決済が終了していませんので、変更する事ができません。"
}

func (e *PG_E11010002) CanRetry() bool {
	return false
}
