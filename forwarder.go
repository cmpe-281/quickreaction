package main

import (
    "bufio"
    "github.com/afex/hystrix-go/hystrix"
    "net/rpc"
    "os"
    "fmt"
)

func main() {
	port1 := os.Args[1]
	port2 := os.Args[2]

	in := bufio.NewReader(os.Stdin)
    for {
        line, _, err := in.ReadLine()
        if err != nil {
            fmt.Println(err)
        }
        output := make(chan bool, 1)
		errors := hystrix.Go("my_command", func() error {
		    _, err := rpc.Dial("tcp", port1)
			if err != nil {
				return err
			}
		    output <- true
		    return nil
		}, nil)

		select {
		case out := <-output:
			fmt.Println("success: ", out)

		    c, e := rpc.Dial("tcp", port1)
			if e != nil {
				panic(e)
			}

			var r int
			e = c.Call("List.IncomingData", string(line), &r)

		case err := <-errors:
			fmt.Println("failure reason:", err)

		    c, e := rpc.Dial("tcp", port2)
			if e != nil {
				panic(e)
			}

			var r int
			e = c.Call("List.IncomingData", string(line), &r)
		}
    }
}