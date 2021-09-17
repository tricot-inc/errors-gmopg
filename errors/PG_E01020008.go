// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01020008 struct{
}

func (e *PG_E01020008) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01020008) Message() string {
    return "ショップパスワードに半角英数字以外の文字が含まれているか、10文字を超えています。"
}

func (e *PG_E01020008) CanRetry() bool {
    return false
}