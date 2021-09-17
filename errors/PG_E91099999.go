// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E91099999 struct{
}

func (e *PG_E91099999) Error() string {
    return "システムエラー(想定外) 購入者に取引失敗を表示し、問い合わせにて状況を確認後、場合によって取消を行ってください。"
}

func (e *PG_E91099999) Message() string {
    return "決済処理に失敗しました。申し訳ございませんが、しばらく時間をあけて購入画面からやり直してください。"
}

func (e *PG_E91099999) CanRetry() bool {
    return false
}