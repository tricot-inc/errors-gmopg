// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01530001 struct{
}

func (e *PG_E01530001) Error() string {
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01530001) Message() string {
    return "決済結果戻り先URLが設定されていません。"
}

func (e *PG_E01530001) CanRetry() bool {
    return false
}