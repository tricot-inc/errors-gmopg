// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01780002 struct{
}

func (e *PG_E01780002) Error() string {
    return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01780002) Message() string {
    return "有効性チェック有無が不正です。"
}

func (e *PG_E01780002) CanRetry() bool {
    return false
}