package generics

import (
	"fmt"
	"time"

	"github.com/samber/lo"
)

func RunLibExample() {
	users := []User{
		{ID: 1, Name: "トム", Age: 30, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "ジェリー", Age: 20, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, Name: "サム", Age: 10, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 4, Name: "サラ", Age: 23, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 5, Name: "ケビン", Age: 40, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 6, Name: "ルーシー", Age: 43, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	newAgeList := lo.Map(users, func(user User, _ int) uint64 {
		return user.Age
	})

	fmt.Println(newAgeList)

	type AgeGroup struct {
		Label  string
		Gender string
	}

	agerMap := lo.GroupBy(users, func(user User) AgeGroup {
		return AgeGroup{Label: fmt.Sprintf("%v代", (user.Age/10)*10), Gender: user.Gender}
	})
	fmt.Println(agerMap)
}
