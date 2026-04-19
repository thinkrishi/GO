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
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the full sentence")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)
	fmt.Println("the sentence is : ", sentence)
}
