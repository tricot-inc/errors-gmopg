// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01510012 struct{
}

func (e *PG_E01510012) Error() string {
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01510012) Message() string {
    return "購買情報文字列が他の項目と矛盾しています。"
}

func (e *PG_E01510012) CanRetry() bool {
    return false
}