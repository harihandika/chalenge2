package main

import "fmt"

func main() {

  var bill int
  var discount int
  var resultdiscount int
  var price int
  var afterDiscount int
  fmt.Println("Enter bill amount:")
  fmt.Scanf("%d", &bill)
  if bill >= 50000 && bill < 70000{
	discount = 25
  } else if bill >= 70000{
	discount = 50
  }else if bill < 50000{
	discount = 0
  }
resultdiscount= (discount / 100)
  fmt.Println(resultdiscount)
 price = (bill * resultdiscount)
  if  discount == 25 && price >= 20000{
	price = 20000
  } else if discount == 50 && price >= 45000{
	price = 45000
  }else{
	price = (bill * discount / 100)
  }
//   fmt.Println("Your discount percentage:")
//   fmt.Scanf("%d", &discount)
  afterDiscount = bill - price
  fmt.Println(price)
  fmt.Println("After discount your bill is: " , afterDiscount)
  }