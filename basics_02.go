// ths program is for thaking a complete inoput of the line like we read in c++ getline(str,cin)
/**import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("full sentence: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("sentence is : ", name)
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a sen:")
	senc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error in reading input")
		return
	}
	senc = strings.TrimSpace(senc)
	fmt.Println("sentence is:", senc, err)

}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a number:")
	senc, _ := reader.ReadString('\n')
	senc = strings.TrimSpace(senc)
	num, err := strconv.Atoi(senc)
	if err != nil {
		fmt.Println("error in reading input")
		return
	}

	fmt.Println("number is:", num)
}
