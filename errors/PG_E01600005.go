// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01600005 struct{
}

func (e *PG_E01600005) Error() string {
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01600005) Message() string {
    return "会員情報チェック文字列が最大文字数を超えています。"
}

func (e *PG_E01600005) CanRetry() bool {
    return false
}