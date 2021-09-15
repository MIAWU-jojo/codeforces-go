// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`"AAB"`}, {`"AAABBC"`}}
	exampleOuts := [][]string{{`8`}, {`188`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, numTilePossibilities, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}