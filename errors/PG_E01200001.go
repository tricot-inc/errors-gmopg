// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01200001 struct {
}

func (e *PG_E01200001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01200001) Message() string {
	return "サイトパスワードが指定されていません。"
}

func (e *PG_E01200001) Code() string {
	return "E01200001"
}

func (e *PG_E01200001) CanRetry() bool {
	return false
}
