// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01790011 struct {
}

func (e *PG_E01790011) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01790011) Message() string {
	return "チェック実施日が最大桁数を超えています。"
}

func (e *PG_E01790011) CanRetry() bool {
	return false
}
