package reflect

import (
	"fmt"
	"log"
	"testing"
)

func TestReflectMap(t *testing.T) {
	sp := New()

	sp.Add("add", func(a,b int) int {return a + b})

	call, err := sp.Call("add", 12, 23)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(call)
	i,b := call[0].Interface().(int)
	if b {
		fmt.Println(i)
	}
}