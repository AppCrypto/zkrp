package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"zkrp/bulletproofs"
)

func main() {

	age := 30
	ageupper := 50
	agelower := 20

	argCount := len(os.Args[1:])

	if argCount > 0 {
		age, _ = strconv.Atoi(os.Args[1])
	}
	if argCount > 1 {
		agelower, _ = strconv.Atoi(os.Args[2])
	}
	if argCount > 2 {
		ageupper, _ = strconv.Atoi(os.Args[3])
	}

	params, _ := bulletproofs.SetupGeneric(int64(agelower), int64(ageupper))

	bigSecret := new(big.Int).SetInt64(int64(age))

	proof, _ := bulletproofs.ProveGeneric(bigSecret, params)

	jsonEncoded, _ := json.Marshal(proof)

	var decodedProof bulletproofs.ProofBPRP
	_ = json.Unmarshal(jsonEncoded, &decodedProof)

	rtn, _ := decodedProof.Verify()

	fmt.Printf("Secret age: %d\n\n", age)
	if rtn {
		fmt.Printf("Age verified to be between %d ad %d\n\n", agelower, ageupper)
	} else {
		fmt.Printf("Age NOT verified to be between %d ad %d\n\n", agelower, ageupper)
	}
	fmt.Printf("Proof P1 A: %s\n\n", proof.P1.A)
	fmt.Printf("Proof P1 Mu: %s\n\n", proof.P1.Mu)
	fmt.Printf("Proof P1 S: %s\n\n", proof.P1.S)
	fmt.Printf("Proof P1 T1: %s\n\n", proof.P1.T1)
	fmt.Printf("Proof P1 T2: %s\n\n", proof.P1.T2)
	fmt.Printf("Proof P1 Taux: %s\n\n", proof.P1.Taux)
	fmt.Printf("Proof P1 Tprime: %s\n\n", proof.P1.Tprime)
	fmt.Printf("Proof P1 V: %s\n\n", proof.P1.V)
	fmt.Printf("Proof P1 A.X: %s\n\n", proof.P1.A.X)
	fmt.Printf("Proof P1 V.X: %s\n\n", proof.P1.V.X)
	fmt.Printf("Proof P1 V.Y: %s\n\n", proof.P1.V.Y)

	fmt.Printf("Proof (showing first 400 characters): %s\n\n", jsonEncoded[:400])
}
