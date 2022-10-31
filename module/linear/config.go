package linear

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/random/crypto"
	"github.com/phoebetron/getlin/serial"
)

type Config struct {
	// Inp is the length of the input vector received. This number is equivalent
	// to the number of Tsetlin Automata (TAs) managed by each and every Clause
	// within this module.
	//
	//     I    [0 1 1 0 1 0 1 1 1]
	//
	Inp int
	// Out is the length of the output vector returned. This number is
	// equivalent to the number of Clauses to be managed by this module.
	//
	//     O          [0 0 1]
	//
	Out int
	// Ran provides randomization primitives for guaranteeing stochastic
	// properties of the internal feedback mechanisms.
	Ran getlin.Random
	// Ser is a serialization method enabling the Clause to be stored and
	// loaded. Algorithms for JSON or Protocl Buffers may be used here.
	Ser serial.Interface
	// Sta is the number of states along a single side of the Tsetlin Automata
	// (TAs) states distribution. See clause.Config for more information.
	Sta int
}

func (c Config) Ensure() Config {
	if c.Ran == nil {
		c.Ran = crypto.New(crypto.Config{})
	}

	return c
}

func (c Config) Verify() {
	if c.Inp == 0 {
		panic("Config.Inp must not be empty")
	}
	if c.Out == 0 {
		panic("Config.Out must not be empty")
	}
	if c.Ran == nil {
		panic("Config.Ran must not be empty")
	}
	// if c.Ser == nil {
	// 	panic("Config.Ser must not be empty")
	// }
	if c.Sta == 0 {
		panic("Config.Sta must not be empty")
	}
}
