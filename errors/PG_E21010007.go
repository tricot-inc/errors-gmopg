// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E21010007 struct {
}

func (e *PG_E21010007) Error() string {
	return "3Dセキュア認証エラー(登録確認) 決済を中止して、取引が終了していない事を通知してください。"
}

func (e *PG_E21010007) Message() string {
	return "3Dセキュア認証に失敗しました。もう一度、購入画面からやり直してください。"
}

func (e *PG_E21010007) CanRetry() bool {
	return false
}
