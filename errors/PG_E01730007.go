// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01730007 struct {
}

func (e *PG_E01730007) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01730007) Message() string {
	return "ボーナス金額に数字以外の文字が含まれています。"
}

func (e *PG_E01730007) CanRetry() bool {
	return false
}
