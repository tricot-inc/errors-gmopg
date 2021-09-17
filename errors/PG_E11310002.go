// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11310002 struct{
}

func (e *PG_E11310002) Error() string{
    return "リンク決済エラー すでに取引が完了している可能性があります。"
}

func (e *PG_E11310002) Message() string{
    return "この取引はリンク決済を実行できません。"
}