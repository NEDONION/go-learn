package main

import (
	"fmt"
	"go-learn/oop/model"
)

func main() {
	/*
			面相对象的三大特征:
				- 封装
				- 集成
				- 多态

		类型可组合
	*/

	lucas := model.User{
		Name:      "Lucas Hu",
		Age:       18,
		Height:    180,
		Education: "Tulane",
		Hobby:     []string{"Hiking", "Coding"},
		MoreInfo: map[string]interface{}{
			"Married": false,
			"Salary":  500,
			"Friends": []string{"Tom", "Jerry"},
		},
	}

	fmt.Println("The user's name is: ", lucas.Name)
	fmt.Println("The user's age is: ", lucas.Age)

}
