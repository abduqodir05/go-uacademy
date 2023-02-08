package main

import "fmt"

type Product struct {
	name  string
	price int
}
type Category struct {
	name     string
	Products []Product
}

func main() {	
	total:=0
	categories := []Category{
		{"mevalar", []Product{
			{"olma", 15000},
			{"nok", 20000},
		},
		},
		{"Ichimliklar", []Product{
			{"cola", 10000},
			{"montella", 5000},
		},
		},
	}

	for i, v := range categories {
		fmt.Printf("%v. %v \n", i+1, v.name)
	}

	p:
	var categories_option int
	fmt.Println("choose any category")
	fmt.Scanf("%v\n", &categories_option)

	a := categories[categories_option-1]
	for i, v := range a.Products {
		fmt.Printf("%v. %v\n",i+1,v.name)
	}
	fmt.Println("-1 . Tamonlash")
	var products_option int
	fmt.Println("choose any product")
	fmt.Scanf("%v\n",&products_option)

	if products_option == -1 {
		fmt.Println(total)
		return
	}
	
	b := a.Products[products_option-1]
	fmt.Println(b)


	var countProduct int
	fmt.Println("how many do you want")
	fmt.Scanf("%v\n",&countProduct)

	total+= b.price*countProduct
	fmt.Println(total)
	goto p

}
