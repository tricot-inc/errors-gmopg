// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01790007 struct{
}

func (e *PG_E01790007) Error() string{
    return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01790007) Message() string{
    return "チェック実施日が不正です。"
}