package main

///////////////// Imports ////////////////////

import (
	"log"
	"net"
	"net/rpc"
	"fmt"
)

////////// User Defined Types For Server //////////

type Service int 

type Arith struct {
	X int
	Y int
}

type Response struct{
	Z int
}

///////////////// Server Logic ////////////////////
func (t *Service) Add(inputs* Arith, response* Response) error{
	log.Printf("[Log] Adding %d and %d\n", inputs.X, inputs.Y)
	response.Z = inputs.X + inputs.Y
	return nil
}

func (t *Service) Subtract(inputs* Arith, response* Response) error{
	log.Printf("[Log] Subtracting %d from %d\n", inputs.Y, inputs.X)

	response.Z = inputs.X - inputs.Y
	return nil
}

func (t *Service) Multiply(inputs* Arith, response* Response) error{
	log.Printf("[Log] Multiplying %d and %d\n", inputs.X, inputs.Y)
	response.Z = inputs.X * inputs.Y
	return nil
}

//go automatically derferences struct points by the way
func (t *Service) Divide(inputs* Arith, response* Response) error{
	if inputs.Y == 0{
        return fmt.Errorf("cannot divide by zero")	
	}

	log.Printf("Dividing %d from %d\n", inputs.Y, inputs.X)
	response.Z = inputs.X / inputs.Y
	return nil
}


///////////////// Server Entry Point ////////////////////

func main() {
	//intialize the service
	Arithmetic := new(Service)
	err := rpc.RegisterName("Arithmetic", Arithmetic)
	if err != nil{
		log.Fatal("[Fatal] Failed to register Server: ", err)
	}

	//initialize the listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("[Fatal] Rpc failed to start: ", err)
	}

	//defer closure until we exit scope
	defer listener.Close()
	log.Println("RPC server up and running on port 8080")

	//Request handler
	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Println("Connection Error: ", err)
			continue
		}

		//Handle client in go routine
		//We want to give each client an ability to connect alongside other clients
		go rpc.ServeConn(conn)
	}
}