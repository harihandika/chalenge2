package main

import "fmt"

var bill int
var discount int
var DiscountPrice int
var afterDiscount int
var voucher string

func DumbWMerchBerkah(){
  discount = 25
  DiscountPrice = (bill*discount/100)
  if  DiscountPrice >= 20000{
    DiscountPrice=20000
  }else{
    DiscountPrice = (bill*discount/100)
  }
}

func DumbWMerchMurahBanget(){
  discount = 50
  DiscountPrice = (bill*discount/100)
  if  DiscountPrice >= 45000{
    DiscountPrice=45000
  }else{
    DiscountPrice = (bill*discount/100)
  }
}

func main() {
  fmt.Println("Enter bill amount:")
  fmt.Scanf("%d", &bill, &voucher)
  if bill >= 50000 && bill < 70000{
    DumbWMerchBerkah()
    fmt.Println("Your Voucher Discount: DumbWMerchBerkah" )   
  } else if bill >= 70000{
	DumbWMerchMurahBanget()
  fmt.Println("Your Voucher Discount: DumbWMerchMurahBanget" )   
  }else if bill < 50000{
	DiscountPrice = 0
  }
  fmt.Println(bill)
  afterDiscount = bill - DiscountPrice
  fmt.Println("your discount:",discount,"%")
  fmt.Println("Discount Price :",DiscountPrice)
  fmt.Println("After discount your bill is: " , afterDiscount)
  }