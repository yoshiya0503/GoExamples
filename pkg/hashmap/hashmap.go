package hashmap

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

type AgeGroup struct {
	Label  string
	Gender string
}

func (us Users) GroupBy() map[AgeGroup]Users {
	// mapのキーに構造体が設定できる。型があれば、slice, map以外すべて
	// 設定できる
	ageUser := make(map[AgeGroup]Users)
	for _, u := range us {
		key := AgeGroup{Label: fmt.Sprintf("%v代", (u.Age/10)*10), Gender: u.Gender}
		ageUser[key] = append(ageUser[key], u)
	}
	return ageUser
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

	fmt.Println("-------------groupBy example---------------")
	groupByAge := users.GroupBy()
	fmt.Println(groupByAge)
}
