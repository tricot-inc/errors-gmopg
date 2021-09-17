// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E21020002 struct{
}

func (e *PG_E21020002) Error() string {
    return "3Dセキュア認証エラー(登録確認) 決済を中止して、取引が終了していない事を通知してください。"
}

func (e *PG_E21020002) Message() string {
    return "3Dセキュア認証がキャンセルされました。もう一度、購入画面からやり直してください。"
}

func (e *PG_E21020002) CanRetry() bool {
    return false
}