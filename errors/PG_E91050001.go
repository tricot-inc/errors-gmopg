// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E91050001 struct {
}

func (e *PG_E91050001) Error() string {
	return "システムエラー(リンクタイプ) テンプレートの設定内容をご確認ください。"
}

func (e *PG_E91050001) Message() string {
	return "決済処理に失敗しました。"
}

func (e *PG_E91050001) CanRetry() bool {
	return false
}
