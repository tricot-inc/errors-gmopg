// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E21030001 struct{
}

func (e *PG_E21030001) Error() string{
    return "3DS2.0認証失敗(認証初期化) 決済を中止して、取引が終了していない事を通知してください。"
}

func (e *PG_E21030001) Message() string{
    return "3Dセキュア認証に失敗しました。もう一度、購入画面からやり直してください。"
}