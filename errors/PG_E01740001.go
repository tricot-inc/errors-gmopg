// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01740001 struct {
}

func (e *PG_E01740001) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01740001) Message() string {
	return "端末処理通番が指定されていません。"
}

func (e *PG_E01740001) CanRetry() bool {
	return false
}
