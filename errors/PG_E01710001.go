// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01710001 struct{
}

func (e *PG_E01710001) Error() string {
    return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01710001) Message() string {
    return "取引区分(継続課金)が設定されていません。"
}

func (e *PG_E01710001) CanRetry() bool {
    return false
}