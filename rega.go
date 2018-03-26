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

func addState(l []*state, s *state, a *state) []*state { //list of pointers to states, single pointer to a state and accept state
	l = append(l, s)

	if s != a && s.symbol == 0 { //deal with e arrows
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}

	}

	return l
}

func pomatch(po string, s string) bool { //find out if pofix regexp matches the string
	ismatch := false
	ponfa := poregtonfa(po)

	current := []*state{} //array of pointers to state. List of states that we are currently in on NFA
	next := []*state{}    // states we can get to

	current = addState(current[:], ponfa.initial, ponfa.accept) //pass current state, pass initial state and accept state. [:] is slice, for when you pass an array and want to change it

	for _, r := range s { //r is rune. loop through s a character at a time
		for _, c := range current { //c is current state we are in. loop through the current states
			if c.symbol == r { // if the sybmol is the same as the one currently reading from s
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{} //move current states to next states. Next array made blank
	}

	for _, c := range current { //loop through the state that you end up in at the very end
		if c == ponfa.accept {
			ismatch = true
			break // from this point is match will always be true
		}
	}

	return ismatch
}

func main() {
	fmt.Println(pomatch("ab.c*|", "cccc")) //take a regexp in postfix and any string. return back true if the string and regexp match
}
