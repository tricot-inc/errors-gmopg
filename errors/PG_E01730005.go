// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01730005 struct {
}

func (e *PG_E01730005) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01730005) Message() string {
	return "ボーナス金額が最大桁数を超えています。"
}

func (e *PG_E01730005) Code() string {
	return "E01730005"
}

func (e *PG_E01730005) CanRetry() bool {
	return false
}
