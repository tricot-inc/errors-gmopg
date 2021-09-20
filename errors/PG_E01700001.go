// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01700001 struct {
}

func (e *PG_E01700001) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01700001) Message() string {
	return "項目数が誤っています。"
}

func (e *PG_E01700001) Code() string {
	return "E01700001"
}

func (e *PG_E01700001) CanRetry() bool {
	return false
}
