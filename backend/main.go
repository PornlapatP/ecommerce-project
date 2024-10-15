package main

import "fmt"

func main() {
	var x int = 10
	var p *int = &x // p เก็บที่อยู่ของ x

	fmt.Println("ค่าของ x:", x)                // แสดงผล 10
	fmt.Println("ที่อยู่ของ x:", &x)           // แสดงที่อยู่ในหน่วยความจำของ x
	fmt.Println("ค่าที่ pointer p ชี้ไป:", *p) // แสดงค่า 10 ซึ่งเป็นค่าของ x ผ่าน pointer

	*p = 20                                 // เปลี่ยนค่าของ x ผ่าน pointer
	fmt.Println("ค่าของ x หลังเปลี่ยน:", x) // แสดงผล 20
}
