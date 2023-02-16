package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	/*[“I”,“am”,“stupid”,“and”,“weak”]
	用 for 循环遍历该数组并修改为
	[“I”,“am”,“smart”,“and”,“strong”]*/
	var testArray = [5]string{"I", "am", "stupid", "and", "weak"}
	for i := 0; i < 5; i++ {
		if testArray[i] == "stupid" {
			testArray[i] = "smart"
		}
		if testArray[i] == "weak" {
			testArray[i] = "strong"
		}
	}
	fmt.Println(testArray)
}
