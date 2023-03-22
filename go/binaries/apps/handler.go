package main

import sdk "github.com/Cloud-Init/go/binaries/pb"

//import mesh "github.com/keti-openfx/openfx/executor/go/mesh"

func Handler(req sdk.Request) string {
	// mesh call
	//
	// functionName := "<FUNCTIONNAME>"
	// input := string(req.Input)
	// result := mesh.MeshCall(functionName, []byte(input))
	// return result

	// single call
	return "[Go] " + string(req.Input)
}
