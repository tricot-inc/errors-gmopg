// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01710002 struct{
}

func (e *PG_E01710002) Error() string {
    return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01710002) Message() string {
    return "指定された取引区分が存在しません。"
}

func (e *PG_E01710002) CanRetry() bool {
    return false
}