// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_EX1000303 struct{
}

func (e *PG_EX1000303) Error() string{
    return "トークンエラー 指定されたトークンの有効期限が切れています。"
}

func (e *PG_EX1000303) Message() string{
    return "決済処理に失敗しました。もう一度カード番号を入力してください。"
}