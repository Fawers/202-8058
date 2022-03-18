package main

import (
	"fmt"

	"github.com/Fawers/battery-info-go/battery"
)

func main() {
	bi, err := battery.NewForDefaultDevice()

	if err != nil {
		switch err := err.(type) {
		case *battery.DefaultDeviceNotFoundError:
			fmt.Println("O sistema não conseguiu encontrar o disposito padrão")

		case *battery.InvalidDeviceError:
			fmt.Printf("O sistema não encontrou o dispositivo %q\n", err.Name)

		case *battery.CommandError:
			fmt.Printf("Erro ao executar %#q:\n%s\n", err.Cmd.String(), err.Err)
		}
	}

	fmt.Println(bi)
}
