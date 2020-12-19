package tornado

import (
	"errors"
	"github.com/mateigraura/utxo-poc/zkp"
	"math"
)

const levels = 2
const zeroval = "BTC_23,624.98"

type Root struct {
	Hash    string
	IsValid bool
}

type Tree struct {
	Levels       int16
	CurrentIndex int16
	Null         string
	RootHistory  map[string]Root
	Leaves       []string
}

func NewTree() *Tree {
	t := Tree{}
	t.Levels = levels
	t.CurrentIndex = 0
	t.Null = zkp.MiMCBn256Decoded(zeroval)
	t.RootHistory = make(map[string]Root)

	for i := 0; i < int(math.Pow(2, float64(t.Levels))); i++ {
		t.Leaves = append(t.Leaves, zkp.MiMCBn256Decoded(t.Null))
	}
	root := t.MakeTree()
	t.RootHistory[root] = Root{root, true}
	return &t
}

func (t *Tree) Insert(leaf string) error {
	if t.CurrentIndex >= int16(math.Pow(2, float64(t.Levels))) {
		return errors.New("tree is full")
	}
	if t.Leaves[t.CurrentIndex] != zkp.MiMCBn256Decoded(t.Null) {
		return errors.New("leaf not nullified. spot already filled")
	}

	t.Leaves[t.CurrentIndex] = zkp.MiMCBn256Decoded(leaf)
	t.CurrentIndex++

	newRoot := t.MakeTree()
	t.RootHistory[leaf] = Root{newRoot, true}
	return nil
}

func (t *Tree) MakeTree() string {
	newLevel := t.Leaves
	for len(newLevel) > 1 {
		var currLevel []string
		for i := 0; i < len(newLevel); i += 2 {
			newNode := zkp.MiMCBn256Decoded(newLevel[i] + newLevel[i+1])
			currLevel = append(currLevel, newNode)
		}
		newLevel = currLevel
	}

	return newLevel[0]
}

func (t *Tree) HasValidRoot(nullifier string) bool {
	if val, ok := t.RootHistory[nullifier]; ok {
		return val.IsValid
	}
	return false
}
