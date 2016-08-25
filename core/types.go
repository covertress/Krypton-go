// Copyright 2014 The go-krypton Authors
// This file is part of the go-krypton library.
//
// The go-krypton library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-krypton library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-krypton library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"math/big"

	"github.com/krypton/go-krypton/accounts"
	"github.com/krypton/go-krypton/core/state"
	"github.com/krypton/go-krypton/core/types"
	"github.com/krypton/go-krypton/core/vm"
	"github.com/krypton/go-krypton/krdb"
	"github.com/krypton/go-krypton/event"
)

// Validator is an interface which defines the standard for block validation.
//
// The validator is responsible for validating incoming block or, if desired,
// validates headers for fast validation.
//
// ValidateBlock validates the given block and should return an error if it
// failed to do so and should be used for "full" validation.
//
// ValidateHeader validates the given header and parent and returns an error
// if it failed to do so.
//
// ValidateStack validates the given statedb and optionally the receipts and
// gas used. The implementor should decide what to do with the given input.
type Validator interface {
	ValidateBlock(block *types.Block) error
	ValidateHeader(header, parent *types.Header, checkPow bool) error
	ValidateState(block, parent *types.Block, state *state.StateDB, receipts types.Receipts, usedGas *big.Int) error
}

// Processor is an interface for processing blocks using a given initial state.
//
// Process takes the block to be processed and the statedb upon which the
// initial state is based. It should return the receipts generated, amount
// of gas used in the process and return an error if any of the internal rules
// failed.
type Processor interface {
	Process(block *types.Block, statedb *state.StateDB) (types.Receipts, vm.Logs, *big.Int, error)
}

// Backend is an interface defining the basic functionality for an operable node
// with all the functionality to be a functional, valid Krypton operator.
//
// TODO Remove this
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *BlockChain
	TxPool() *TxPool
	ChainDb() krdb.Database
	DappDb() krdb.Database
	EventMux() *event.TypeMux
}
