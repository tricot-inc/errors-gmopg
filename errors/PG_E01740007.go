// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01740007 struct {
}

func (e *PG_E01740007) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01740007) Message() string {
	return "端末処理通番に数字以外の文字が含まれています。"
}

func (e *PG_E01740007) CanRetry() bool {
	return false
}
