// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_42C780000 struct{
}

func (e *PG_42C780000) Error() string{
    return "決通信エラー(CAFISまたはカード会社障害) カード所有者に取引失敗を表示し、問い合わせにて状況を確認してください。"
}

func (e *PG_42C780000) Message() string{
    return "決済処理に失敗しました。申し訳ございませんが、しばらく時間をあけて購入画面からやり直してください。"
}