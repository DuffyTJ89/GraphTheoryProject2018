package main

import (
	"fmt"
)

type state struct {
	//store letter with state. Value is 0
	symbol rune
	//the two arrows that come from the state. pointers to other states
	edge1 *state
	edge2 *state
}

type nfa struct { //keep track of initial and accept states
	initial *state
	accept  *state
}

func poregtonfa(pofix string) *nfa { //return pointer to nfa
	nfaStack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1] // : means give me everything off the statck up to the last element but not including it
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial //join together and push the conceited fragment back to the nfa stack

			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept}) //pops 2 fragments off the stack, joins the accept state of frag1 to initial state of frag2, push the new fragment onto the nfa stack

		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1] //: means give me everything off the statck up to the last element but not including it
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		case '*':
			//only pop one frag off the nfa stack
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//new frag is the old frag with 2 extra states
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept}) //push new frag to nfa stack

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept} //label new accept state with symbol r

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}

	return nfaStack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
