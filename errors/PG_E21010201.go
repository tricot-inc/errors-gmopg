// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E21010201 struct{
}

func (e *PG_E21010201) Error() string {
    return "3Dセキュア必須化エラー(未対象) 決済を中止して、取引が終了していない事を通知してください。"
}

func (e *PG_E21010201) Message() string {
    return "このカードでは取引をする事ができません。3Dセキュア認証に対応したカードをお使いください。"
}

func (e *PG_E21010201) CanRetry() bool {
    return false
}