// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01210002 struct{
}

func (e *PG_E01210002) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01210002) Message() string {
    return "指定されたIDとパスワードのサイトが存在しません。"
}

func (e *PG_E01210002) CanRetry() bool {
    return false
}