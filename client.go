package main

import (
	"log"
	"net/rpc"
)

type Arith struct{
	X int
	Y int
}

type Response struct{
	Z int
}

func main(){
	//initializing the client
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil{
		log.Fatal("[Fatal] Error dialing: ", err)
	}
	
	defer client.Close()

	//arguments and responses
	args := &Arith{10,10}
	addResponse := &Response{}
	mulResponse := &Response{}
	divResponse := &Response{}
	subResponse := &Response{}

	//testing addition
	err = client.Call("Arithmetic.Add", args, addResponse)
	if err != nil{
		log.Fatal("[Fatal] Issue addin ", err)
	}
	log.Printf("Output of Addition: %d", addResponse.Z)

	//testing multiplication
	err = client.Call("Arithmetic.Multiply", args, mulResponse)
	if err != nil{
		log.Fatal("[Fatal] Issue multiplying", err)
	}
	log.Printf("Output of Multiplying: %d", mulResponse.Z)

	//testing division
	err = client.Call("Arithmetic.Divide", args, divResponse)
	if err != nil{
		log.Fatal("[Fatal] Issue dividing", err)
	}
	log.Printf("Output of Division: %d", divResponse.Z)

	//testing subtraction
	err = client.Call("Arithmetic.Subtract", args, subResponse)
	if err != nil{
		log.Fatal("[Fatal] Issue subtracting", err)
	}
	log.Printf("Output of subtraction: %d", subResponse.Z)
}