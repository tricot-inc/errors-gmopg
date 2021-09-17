// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11310004 struct{
}

func (e *PG_E11310004) Error() string{
    return "リンク決済エラー セッションタイムアウト"
}

func (e *PG_E11310004) Message() string{
    return "この取引はリンク決済を実行できません。"
}