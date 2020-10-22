package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/csv"
)

// user defined structs
type problem struct {
	q string
	a string
}

/*
	Functions from down here
*/
func exit(s string, retcode int) {
	fmt.Println(s)
	os.Exit(retcode)
}

func main() {
	filename := flag.String("csv", "problems.csv", "csv file name containing quiz in the format (question, answer)")
	flag.Parse()

	fp,err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Couldn't open file for reading: %s", *filename), 1)
	} else {
		//fp.Close()
		reader := csv.NewReader(fp)
		recs, err := reader.ReadAll()
		if err != nil {
			exit(fmt.Sprintf("Couldn't read from the opened file :%s", *filename), 1)
		}
//		fmt.Println(recs)
		probs := populate(recs)
		play(probs)
	}
}

func populate(recs [][]string) []problem {
	ret := make([]problem, len(recs))
	for i,line := range recs {
		ret[i] = problem {
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func play(probs []problem) {
	score := 0
	for i, prob := range probs {
		fmt.Printf("Q%d. %s = ", i+1, prob.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == prob.a {
			score++
		}
	}
	exit(fmt.Sprintf("Your final score is: %d\tPercentage: %.2f%%", score, float64(100 * score/len(probs))), 0)
}
