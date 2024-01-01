package iteraciones

import (
	"fmt"
)

func Iterar() {

	for i := 0; i <= 100; i += 2 {
		fmt.Println(i)
		i++
	}
}
