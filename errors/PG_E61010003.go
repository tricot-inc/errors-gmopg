// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E61010003 struct {
}

func (e *PG_E61010003) Error() string {
	return "加盟店設定エラー 決済を中止して、問い合わせにて設定状況を確認してください。"
}

func (e *PG_E61010003) Message() string {
	return "決済処理に失敗しました。申し訳ございませんが、しばらく時間をあけて購入画面からやり直してください。"
}

func (e *PG_E61010003) Code() string {
	return "E61010003"
}

func (e *PG_E61010003) CanRetry() bool {
	return false
}
