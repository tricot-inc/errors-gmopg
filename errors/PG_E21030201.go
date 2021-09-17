// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E21030201 struct {
}

func (e *PG_E21030201) Error() string {
	return "3DS2.0認証非対応エラー 決済を中止して、取引が終了していない事を通知してください。"
}

func (e *PG_E21030201) Message() string {
	return "このカードでは取引をする事ができません。3Dセキュア認証に対応したカードをお使いください。"
}

func (e *PG_E21030201) CanRetry() bool {
	return false
}
