package collection

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Age       uint64    `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

func (us Users) FilterBy(condition func(u User) bool) Users {
	// 自分で設定した型にはメソッドを生やすことができる。
	// スライス型にももちろんメソッドが定義できる
	users := make(Users, 0)
	for _, u := range us {
		if condition(u) {
			users = append(users, u)
		}
	}
	return users
}

func RunExample() {
	users := Users{
		{ID: 1, Name: "トム", Age: 30, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "ジェリー", Age: 20, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, Name: "サム", Age: 10, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 4, Name: "サラ", Age: 23, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 5, Name: "ケビン", Age: 40, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 6, Name: "ルーシー", Age: 43, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	teenAger := users.FilterBy(func(u User) bool {
		return u.Age < 20
	})

	fmt.Println("-------------teenAger example---------------")
	fmt.Println(teenAger)
}
