package execution

import "fmt"

type Execution interface {
	Start(tomeePath string) error
	Stop(tomeePath string) error
	Restart(tomeePath string)
}

func Restart(e Execution, tomeePath string) {
	err := e.Stop(tomeePath)
	if err != nil {
		fmt.Println("TomEE isn'type started...")
	}

	fmt.Printf("Starting server..")
	err = e.Start(tomeePath)
	if err != nil {
		fmt.Println("Error during start")
		return
	}

	fmt.Println("TomEE started")
}
