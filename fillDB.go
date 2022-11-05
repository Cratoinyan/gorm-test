package main

func FillDB() {

	db, err := GetDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	//user records
	db.Create(&User{Name: "UmutÇ", Mail: "umutciloglu@gmail.com"})
	db.Create(&User{Name: "Umutİ", Mail: "umutinceer@gmail.com"})
	db.Create(&User{Name: "UmutG", Mail: "umutgerçek@gmail.com"})
	db.Create(&User{Name: "ÖncümKY", Mail: "öncümkorkmazyılmaz@gmail.com"})
	db.Create(&User{Name: "Barışİ", Mail: "barışikinci@gmail.com"})

	//credit card records
	db.Create(&CreditCard{UserID: 1, CardNumber: 1234567891234567, Name: "umut çiloğlu", ExpMonth: 12, ExpYeat: 24, SecCode: 123})
	db.Create(&CreditCard{UserID: 1, CardNumber: 1234567891234567, Name: "umut çiloğlu", ExpMonth: 11, ExpYeat: 27, SecCode: 321})
	db.Create(&CreditCard{UserID: 2, CardNumber: 2234567891234567, Name: "umut inceer", ExpMonth: 12, ExpYeat: 26, SecCode: 281})
	db.Create(&CreditCard{UserID: 2, CardNumber: 2234567891234567, Name: "umut inceer", ExpMonth: 10, ExpYeat: 23, SecCode: 423})
	db.Create(&CreditCard{UserID: 3, CardNumber: 3234567891234567, Name: "umut gerçek", ExpMonth: 12, ExpYeat: 27, SecCode: 523})
	db.Create(&CreditCard{UserID: 3, CardNumber: 3234567891234567, Name: "umut gerçek", ExpMonth: 9, ExpYeat: 28, SecCode: 623})
	db.Create(&CreditCard{UserID: 4, CardNumber: 4234567891234567, Name: "öncüm korkmaz yılmaz", ExpMonth: 12, ExpYeat: 29, SecCode: 153})
	db.Create(&CreditCard{UserID: 4, CardNumber: 4234567891234567, Name: "öncüm korkmaz yılmaz", ExpMonth: 8, ExpYeat: 25, SecCode: 133})
	db.Create(&CreditCard{UserID: 5, CardNumber: 5234567891234567, Name: "barış ikinci", ExpMonth: 12, ExpYeat: 27, SecCode: 103})
	db.Create(&CreditCard{UserID: 5, CardNumber: 5262767891234567, Name: "barış ikinci", ExpMonth: 7, ExpYeat: 24, SecCode: 124})

}
