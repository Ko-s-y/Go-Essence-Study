package main

import (
	"fmt"
	"log"
)

// User 構造体を定義
type User struct {
	Name string
}

// users リストを仮想的に定義
var users = []*User{
	{Name: "Alice"},
	{Name: "Bob"},
	{Name: "Charlie"},
}

func main() {
	// FindUser 関数を使ってユーザー "Bob" を検索し、エラー処理を行う
	user, err := FindUser("Boba")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)
}

// FindUser 関数：指定された名前のユーザーを検索
func FindUser(name string) (*User, error) {
	// findUserFromList 関数を呼び出してユーザーを検索し、エラー処理を行う
	user, err := findUserFromList(users, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// findUserFromList 関数：ユーザーをリストから検索
func findUserFromList(userList []*User, name string) (*User, error) {
	// リスト内のユーザーをループで走査して指定された名前のユーザーを検索
	for _, user := range userList {
		if user.Name == name {
			return user, nil
		}
	}
	// ユーザーが見つからなかった場合、エラーメッセージを返す
	return nil, fmt.Errorf("ユーザー '%s' は見つかりませんでした", name)
}
