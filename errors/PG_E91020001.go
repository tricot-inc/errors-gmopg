// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E91020001 struct{
}

func (e *PG_E91020001) Error() string {
    return "システムエラー(通信) 購入者に取引失敗を表示し、問い合わせにて状況を確認してください。"
}

func (e *PG_E91020001) Message() string {
    return "通信タイムアウトが発生しました。申し訳ございませんが、しばらく時間をあけて購入画面からやり直してください。"
}

func (e *PG_E91020001) CanRetry() bool {
    return false
}