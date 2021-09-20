// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01770002 struct {
}

func (e *PG_E01770002) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01770002) Message() string {
	return "区分が不正です。"
}

func (e *PG_E01770002) Code() string {
	return "E01770002"
}

func (e *PG_E01770002) CanRetry() bool {
	return false
}
