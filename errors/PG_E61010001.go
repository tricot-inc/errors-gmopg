// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E61010001 struct{
}

func (e *PG_E61010001) Error() string{
    return "加盟店設定エラー 決済を中止して、問い合わせにてSPID設定状況を確認してください。"
}

func (e *PG_E61010001) Message() string{
    return "決済処理に失敗しました。申し訳ございませんが、しばらく時間をあけて購入画面からやり直してください。"
}