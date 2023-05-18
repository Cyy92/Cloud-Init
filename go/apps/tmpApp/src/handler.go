package main

import sdk "github.com/Cyy92/Cloud-Init/pkg/go/pb"

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
