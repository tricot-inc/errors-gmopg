// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01740005 struct{
}

func (e *PG_E01740005) Error() string {
    return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01740005) Message() string {
    return "端末処理通番が最大桁数を超えています。"
}

func (e *PG_E01740005) CanRetry() bool {
    return false
}