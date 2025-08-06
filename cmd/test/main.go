package main

import "fmt"

type Product struct {
	Name string
	Price float64
}


func New(pro Product) *Product {
	return &Product{
		Name: pro.Name,
		Price: pro.Price,
	}
}

func (p Product) GetPriceWithTax(taxRate float64) float64 {
	return p.Price * (1 + taxRate)
}


// package main

// import "fmt"
func main(){
	inventory:= make(map[string]int)
	inventory["apple"] = 10
	inventory["banana"] = 5
	inventory["orange"] = 0;


	for k, v := range inventory {
		if v == 0 {
			println(fmt.Sprintf("[%s] is out of stock ", k))
		}else{
			println(fmt.Sprintf("[%s] has [%d] ", k, v))
		}
	}
}


func CalculateAverage(s []int) (float64, error) {
	if(len(s) == 0) {
		return float64(0), fmt.Errorf("cannot calculate average of an empty slice")
	}
	add:=0

	for _, v:= range s {
		add += v
	}
	avg:= float64(add) / float64(len(s))
	return avg, nil
}