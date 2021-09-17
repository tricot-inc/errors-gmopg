// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010001 struct {
}

func (e *PG_E11010001) Error() string {
	return "取引エラー 決済を中止して、取引ができない事を通知してください。"
}

func (e *PG_E11010001) Message() string {
	return "この取引はすでに決済が終了しています。"
}

func (e *PG_E11010001) Code() string {
	return "E11010001"
}

func (e *PG_E11010001) CanRetry() bool {
	return false
}
