package pkg

import (
	"fmt"
	"os"
)

func Register()  {
	fmt.Println("PKG IS REGISTERED!")
	os.Exit(1)
}
