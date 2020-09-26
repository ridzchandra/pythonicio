package pythonicio

import (
	"bufio"
	"fmt"
	"os"
)

// Input ... Python like console input
func Input(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}
