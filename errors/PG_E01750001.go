// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01750001 struct {
}

func (e *PG_E01750001) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01750001) Message() string {
	return "利用日が指定されていません。"
}

func (e *PG_E01750001) CanRetry() bool {
	return false
}
