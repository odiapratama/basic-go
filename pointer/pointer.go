package main

import "fmt"

// Address struct
type Address struct {
	City, Province, Region string
}

func changeRegionAddress(address *Address, city string) {
	address.Region = city
}

func pointer() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1

	fmt.Println(address1)
	fmt.Println(address2)

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	address3 := new(Address)
	address3.City = "Cilegon"
	address3.Province = "Banten"
	address3.Region = "Indonesia"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := &Address{"Denpasar", "Bali", ""}
	fmt.Println(address4)

	changeRegionAddress(address4, "Indonesia")
	fmt.Println(address4)
}
