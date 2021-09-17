// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01750008 struct{
}

func (e *PG_E01750008) Error() string {
    return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01750008) Message() string {
    return "利用日の書式が正しくありません。"
}

func (e *PG_E01750008) CanRetry() bool {
    return false
}