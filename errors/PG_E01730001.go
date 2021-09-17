// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01730001 struct{
}

func (e *PG_E01730001) Error() string{
    return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01730001) Message() string{
    return "ボーナス金額が指定されていません。"
}