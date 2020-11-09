package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/csv"
	"time"
	"math/rand"
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
	timelimit := flag.Int("time", 30, "Time limit within which the answer has to be entered")
	shuffle := flag.Bool("shuf", false, "Questions come in random order (may repeat)")
	flag.Parse()

	fp,err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Couldn't open file for reading: %s", *filename), 1)
	} else {
		defer fp.Close()
		reader := csv.NewReader(fp)
		recs, err := reader.ReadAll()
		if err != nil {
			exit(fmt.Sprintf("Couldn't read from the opened file :%s", *filename), 1)
		}
//		fmt.Println(recs)
		probs := populate(recs)
		play(probs, *timelimit, *shuffle)
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

func play(probs []problem, limit int, shuffle bool) {
	score := 0
	timer := time.NewTimer(time.Duration(limit) * time.Second)

	if shuffle == false {
problemloop_if:
		for i, prob := range probs {
			fmt.Printf("Q%d. %s = ", i+1, prob.q)
			timer.Reset(time.Duration(limit) * time.Second)
			anschan := make(chan string)
			go func() {
				var answer string
				fmt.Scanf("%s\n", &answer)
				anschan <- answer
			}() // anonymous go routine
			select {
				case <-timer.C:
					fmt.Println()
					break problemloop_if
				case answer := <- anschan:
					if answer == prob.a {
						score++
					}
					if !timer.Stop() {
						<- timer.C
					}	
			}
		}
	} else {
		rand.Seed(time.Now().UnixNano())
		totprobs := len(probs)
problemloop_else:
		for i := 0; i < totprobs; i++ {
			prob := probs[rand.Intn(totprobs)]
			fmt.Printf("Q%d. %s = ", i+1, prob.q)
            timer.Reset(time.Duration(limit) * time.Second)
            anschan := make(chan string)
            go func() {
                var answer string
                fmt.Scanf("%s\n", &answer)
                anschan <- answer
            }() // anonymous go routine
            select {
                case <-timer.C:
                    fmt.Println()
                    break problemloop_else
                case answer := <- anschan:
                    if answer == prob.a {
                        score++
                    }
                    if !timer.Stop() {
                        <- timer.C
                    }
            }
		}
	}
	exit(fmt.Sprintf("Your final score is: %d\tPercentage: %.2f%%", score, float64(100 * score/len(probs))), 0)
}

/* func (prob *[]problem) que() string {
	rand.Seed(time.Now().UnixNano())
	return prob[rand.Intn(len(prob))]
}	*/
