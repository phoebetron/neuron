package mapper

import "github.com/phoebetron/getlin"

type Linker struct {
	// Abo contains all Modules within the previous layer.
	Abo []getlin.Module
	// Bel contains all Modules within the next layer.
	Bel []getlin.Module
	// Ind is the index of this Linker's Module within the current graph.
	Ind int
	// Lay is the index of the layer this Linker's Module resides in.
	Lay int
	// Tru is the true label index range assigned to the Module linked here.
	Tru [2]int
}
