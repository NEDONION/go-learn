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

	ned := model.NewUser("Ned",
		20,
		190,
		"SCU",
		[]string{"Outdoors"},
		nil)

	fmt.Println("The user's name is: ", lucas.Name)
	fmt.Println("The user's age is: ", lucas.Age)
	fmt.Printf("Ned =%v\n", ned)

	coke := model.Product{}
	coke.SetName("Coke")
	coke.SetPrice(3)
	fmt.Printf("coke =%v\n", coke)

	paymentArgs := model.PaymentArgs{
		AppID:       "super-AppID",
		MerchantID:  "super-MerchantID",
		Key:         "super-Key",
		CallBackUrl: "api.payment.com/super",
	}

	alipay := model.Alipay{
		PaymentArgs: model.PaymentArgs{
			AppID:       "alipay-AppID",
			MerchantID:  "alipay-MerchantID",
			Key:         "alipay-Key",
			CallBackUrl: "api.payment.com/alipay",
		},
		AlipayOpenID: "AlipayOpenID",
	}

	paymentArgs.Info() // Info = &{super-AppID super-MerchantID super-Key api.payment.com/super}
	alipay.Info()      // Info = &{alipay-AppID alipay-MerchantID alipay-Key api.payment.com/alipay}

}
