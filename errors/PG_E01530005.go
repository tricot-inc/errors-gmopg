// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01530005 struct{
}

func (e *PG_E01530005) Error() string {
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01530005) Message() string {
    return "決済結果戻り先URLが最大文字数を超えています。"
}

func (e *PG_E01530005) CanRetry() bool {
    return false
}