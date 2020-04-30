package main

import "fmt"

//RemoveDuplicates removes duplicates
func RemoveDuplicates(toEmail string, bcc []string) []string {

	for i, email := range bcc {
		if toEmail == email {
			fmt.Printf("The recipient's email :%v, is duplicated.", toEmail)
			bcc[i] = bcc[len(bcc)-1]
			bcc[len(bcc)-1] = ""
			bcc = bcc[:len(bcc)-1]

		}

	}
	fmt.Printf(" Here's the new BCC list: %v", bcc)
	return bcc
}

func main() {
	RemoveDuplicates("test@gmail.com",
		[]string{"test@test.com", "test@gmail.com", "didisaytest@gmail.com"})

}
