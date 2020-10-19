package natas


import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Printf("%-8s | %-10s | %-8s | %-0s\n", "Puzzle", "Solved", "Username", "Password")
	fmt.Println("===================================================================")
	passwords := []string{"natas0"}
	Init()
	for i, solution := range(Solutions){
		fmt.Printf("%-8s | %-10s | %-8s | %-0s", "Natas " + strconv.Itoa(i), "Not Solved", "natas" + strconv.Itoa(i), "")
		newPass := solution("natas" + strconv.Itoa(i), passwords[i])
		passwords = append(passwords, newPass)
		fmt.Printf("\r%-8s | %-10s | %-8s | %-0s\n", "Natas " + strconv.Itoa(i), "Solved", "natas" + strconv.Itoa(i), newPass)
	}

}
