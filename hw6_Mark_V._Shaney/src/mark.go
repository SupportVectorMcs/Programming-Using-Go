/** Name: Long He
* Andrew ID: longh
* date: 10/27/2014
* version: 1.0
*/

// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Generating random text: a Markov chain algorithm

Based on the program presented in the "Design and Implementation" chapter
of The Practice of Programming (Kernighan and Pike, Addison-Wesley 1999).
See also Computer Recreations, Scientific American 260, 122 - 125 (1989).

A Markov chain algorithm generates text by creating a statistical model of
potential textual suffixes for a given prefix. Consider this text:

	I am not a number! I am a free man!

Our Markov chain algorithm would arrange this text into this set of prefixes
and suffixes, or "chain": (This table assumes a prefix length of two words.)

	Prefix       Suffix

	"" ""        I
	"" I         am
	I am         a
	I am         not
	a free       man!
	am a         free
	am not       a
	a number!    I
	number! I    am
	not a        number!

To generate text using this table we select an initial prefix ("I am", for
example), choose one of the suffixes associated with that prefix at random
with probability determined by the input statistics ("a"),
and then create a new prefix by removing the first word from the prefix
and appending the suffix (making the new prefix is "am a"). Repeat this process
until we can't find any suffixes for the current prefix or we exceed the word
limit. (The word limit is necessary as the chain table may contain cycles.)

Our version of this program reads text from standard input, parsing it into a
Markov chain, and writes generated text to standard output.
The prefix and output lengths can be specified using the -prefix and -words
flags on the command-line.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
	"time"
)

// Prefix is a Markov chain prefix of one or more words.
type Prefix []string

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

// BuildFreqMap takes the modelLine and prefixLen and put the prefixes and
// suffixes into a chain based on the given modelFile
func (c *Chain) BuildFreqMap(modelLine []string, prefixLen int) {
	var prefixes []string
	prefixes = make([]string,0)
	// for each line of the modelLine
	for i := 1; i < len(modelLine) - 1; i++ {
		// split this line into words
		words := strings.Split(modelLine[i], " ")
		var temp []string
		temp = make([]string,0)
		// get the prefix from these words
		for j := 0; j < prefixLen; j++ {
			temp = append(temp, words[j])
		}
		prefixes = append(prefixes, strings.Join(temp, " "))
	}
	// remove "" from prefixes
	for i := 0; i < len(prefixes); i++ {
		prefixes[i] = strings.Replace(prefixes[i], "\"\"", "", -1)
	}

	var suffixes [][]string
	suffixes = make([][]string,len(modelLine)-2)
	// for each line of the modelLine
	for i := 1; i < len(modelLine) - 1; i++ {
		// split this line into words
		words := strings.Split(modelLine[i], " ")
		var temp []string
		temp = make([]string,0)
		// get the suffixes from these words
		for j := prefixLen; j < len(words) - 1; j++ {
			temp = append(temp, words[j])
			j++
		}
		// get the same suffixes from these words
		for j := prefixLen + 1; j < len(words); j++ {
			a,_ := strconv.Atoi(words[j])
			for k := 0; k < a - 1; k++ {
				temp = append(temp, words[j - 1])
			}
			j++
		}
		// combine all suffixes
		suffixes[i-1] = temp
	}
	// put prefixes and suffixes into the chain
	for i := 0; i < len(modelLine) - 2; i++ {
		c.chain[prefixes[i]] = suffixes[i]
	}
}

// WriteTable takes the modelName and prefixLen and 
// writes the frequency table into a file.
func (c *Chain) WriteTable(modelName string, prefixLen int) {
	// Calculate the frequency
	var freqMap map[string]map[string]int
	freqMap = make(map[string]map[string]int)

	// Save all the prefixes into the freqMap
	for prefix,_ := range c.chain {
		freqMap[prefix] = make(map[string]int)
	}

	// Save all the suffixes into the freqMap
	for prefix,suffixSet := range c.chain {
		for _,eachSuffix := range suffixSet {
			// if there are same suffixes
			if freqMap[prefix][eachSuffix] >= 1 {
				freqMap[prefix][eachSuffix] += 1
			} else {
				freqMap[prefix][eachSuffix] = 1
			}
		}
	}

	// Open files for writing
	outFile, err := os.Create(modelName)
    if err != nil {
        fmt.Println("Sorry: couldn’t create the file!")
    }

    // Write the first line
    fmt.Fprintf(outFile, "%d\n", prefixLen)

    // Write the following lines
    for prefix,suffixMap := range freqMap {
    	if strings.TrimPrefix(prefix, " ") != prefix {
    		// write "" for different cases
    		for strings.TrimPrefix(prefix, " ") != prefix {
    			fmt.Fprintf(outFile, "\"\" ")
    			prefix = strings.TrimPrefix(prefix, " ")
    		}
    		if prefix == "" {
    			fmt.Fprintf(outFile, "\"\"")
    		}
    		fmt.Fprintf(outFile, "%v ", prefix)
    	} else {
    		fmt.Fprintf(outFile, "%v ", prefix)
    	}
		for suffixes,freq := range suffixMap {
			fmt.Fprintf(outFile, "%v ", suffixes)
			fmt.Fprintf(outFile, "%d ", freq)
		}
		fmt.Fprintln(outFile)
	}
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}

func main() {
	// "read" command
	if os.Args[1] == "read" {
		// make sure number of files is valid
    	prefixLen, err := strconv.Atoi(os.Args[2])
    	if err != nil {
        fmt.Println("Error: prefixLen must be an integer.")
        return
    	} else if prefixLen < 0 {
        	fmt.Println("Error: prefixLen must be positive.")
        	return
    	}

    	// name of the model
    	modelName := os.Args[3]

    	// Initialize a new Chain.
		c := NewChain(prefixLen)

		// Build reads text from the provided Reader and
		// parses it into prefixes and suffixes that are stored in Chain.
    	for i := 0; i < len(os.Args) - 4; i++ {
    		file, err := os.Open(os.Args[i + 4])
			if err != nil {
				fmt.Println("Error: something went wrong opening the file.")
				fmt.Println("Probably you gave the wrong filename.")
				os.Exit(3)
			}
			c.Build(file)
    	}		
    	// Use the given length of prefix and name of model to
    	// write the frequency table into a file
    	c.WriteTable(modelName, prefixLen)
	} else if os.Args[1] == "generate" { // "generate" command
		// make sure number of words is valid
    	numWords, err := strconv.Atoi(os.Args[3])
    	if err != nil {
        fmt.Println("Error: numWords must be an integer.")
        return
    	} else if numWords < 0 {
        	fmt.Println("Error: numWords must be positive.")
        	return
    	}

    	// Seed the random number generator.
    	rand.Seed(time.Now().UnixNano()) 

    	// read the file into a string
    	s,_ := ioutil.ReadFile(os.Args[2])
    	modelString := string(s)

    	// split the model file string line by line
    	modelLine := strings.Split(modelString, "\n")

    	// get the length of the prefix
    	prefixLen,_ := strconv.Atoi(modelLine[0])

    	// Initialize a new Chain.
    	c := NewChain(prefixLen)

    	// Use the given converted modelFile and prefixLen to build 
    	// a frequncy table
    	c.BuildFreqMap(modelLine, prefixLen)
    	// Generate text based on the given number of words.
    	text := c.Generate(numWords) 
    	// Write text to standard output.
    	fmt.Println(text)            
	}
}