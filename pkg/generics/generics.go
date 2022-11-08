package generics

import (
	"fmt"
	"time"

	"github.com/samber/lo"
)

func Map[T, U any](collection []T, fn func(T) U) []U {
	var result []U
	for _, item := range collection {
		result = append(result, fn(item))
	}
	return result
}

func Filter[T any](collection []T, fn func(T) bool) []T {
	var result []T
	for _, item := range collection {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Age       uint64    `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func RunExample() {
	users := []User{
		{ID: 1, Name: "トム", Age: 30, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 2, Name: "ジェリー", Age: 20, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 3, Name: "サム", Age: 10, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 4, Name: "サラ", Age: 23, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 5, Name: "ケビン", Age: 40, Gender: "男性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: 6, Name: "ルーシー", Age: 43, Gender: "女性", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	ageList := Map(users, func(user User) uint64 {
		return user.Age
	})
	fmt.Println(ageList)

	kidsList := Filter(users, func(user User) bool {
		return user.Age < 20
	})
	fmt.Println(kidsList)

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
	/*
	  nameLength := lo.GroupBy(users, func(user User) int {
	    return len(user.Name)
	    })
	  fmt.Println(nameLength)
	*/
}
