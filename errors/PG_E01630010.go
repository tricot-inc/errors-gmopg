// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01630010 struct{
}

func (e *PG_E01630010) Error() string {
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01630010) Message() string {
    return "取引後カード登録時、取引の会員IDとパラメータの会員IDが一致しません。"
}

func (e *PG_E01630010) CanRetry() bool {
    return false
}