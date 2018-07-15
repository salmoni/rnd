package main

/*
rnd - a command line program to generate a list of random numbers

rnd [parameters] type distribution number min/stddev max/mean

type = type of data: 'f' for float (64 bit), 'i' for integers
distribution = 'u' = uniform distribution, 'n' = normal distribution
number = how many numbers to produce (must be int)
min/stddev = either minimum value (for uniform distributions) or standard deviation (for normal distributions)
max/mean = either maximum value (for uniform distributions) or mean (for normal distributions)

(c) 2018 Alan James Salmoni. Released under the MIT licence

*/

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var dataFloat []float64
	//var dataInt []int

	programArgs := os.Args
	fmt.Println("Length of args = ", len(programArgs), programArgs[2])
	if len(programArgs) < 3 {
		// not enough parameters
		notEnoughParametersError()
	}
	seqNumber, error := strconv.Atoi(programArgs[3])
	if error != nil {
		incorrectNumberError()
	}

	seqDataType := programArgs[1]
	seqDataDist := programArgs[2]
	fmt.Println("DEBUG: ", seqDataType, seqDataDist, seqNumber)

	rand.Seed(574389)

	if seqDataType == "f" {
		// generate a list of floats
		fmt.Println("DEBUG: Floats")

		if len(programArgs) == 4 {
			fmt.Println("Doing floats with 3 params")
			if seqDataDist == "u" {
				// Uniform distribution of floats
				for idx := 0; idx < seqNumber; idx++ {
					value := rand.Float64()
					dataFloat = append(dataFloat, value)
				}
				printFloat(dataFloat)
			} else if seqDataDist == "n" {
				// normal distribution of floats
				for idx := 0; idx < seqNumber; idx++ {
					value := rand.NormFloat64()
					dataFloat = append(dataFloat, value)
				}
				printFloat(dataFloat)
			} else {
				// Error - didn't specify distribution
				incorrectDistributionError()
			}
		} else if len(programArgs) == 6 {
			seqMin, error := strconv.ParseFloat(programArgs[3], 64)
			if error != nil {
				cannotConvertParameterError()
			}
			seqMax, error := strconv.ParseFloat(programArgs[4], 64)
			if error != nil {
				cannotConvertParameterError()
			}
			seqRange := seqMax - seqMin
			if seqRange < 0 {
				if error != nil {
					cannotConvertParameterError()
				}
			}
			if seqDataDist == "u" {
				// Uniform distribution of floats
				for idx := 0; idx < seqNumber; idx++ {
					value := (rand.Float64() * seqRange) + seqMin
					dataFloat = append(dataFloat, value)
				}
				printFloat(dataFloat)
			} else if seqDataDist == "n" {
				// normal distribution of floats
				for idx := 0; idx < seqNumber; idx++ {
					value := (rand.NormFloat64() * seqRange) + seqMin
					dataFloat = append(dataFloat, value)
				}
				printFloat(dataFloat)
			} else {
				// Error - didn't specify distribution
				incorrectDistributionError()
			}
		}
	} else {
		fmt.Println("DEBUG: Ints")
		// cannot understand
		generalError()
	}
}

func printFloat(data []float64) {
	// Print out floats
	n := len(data)
	if n < 1 {
		os.Exit(0)
	} else if n == 1 {
		// print out one
		fmt.Println(data[0])
	} else {
		outstr := strconv.FormatFloat(data[0], 'f', -1, 64)
		for idx := 1; idx < n; idx++ {
			outstr = outstr + ", " + strconv.FormatFloat(data[idx], 'f', -1, 64)
		}
		fmt.Println(outstr)
	}

	os.Exit(0)
}

func notEnoughParametersError() {
	fmt.Println("Error: Not enough parameters")
	os.Exit(3)
}

func generalError() {
	fmt.Println("General error")
	os.Exit(3)
}

func incorrectNumberError() {
	fmt.Println("Error: Number not specified properly")
	os.Exit(3)
}

func incorrectDistributionError() {
	fmt.Println("Error: Need to specify uniform or normal distribution")
	os.Exit(3)
}

func cannotConvertParameterError() {
	fmt.Println("Error: Cannot convert parameter (min or max)")
	os.Exit(3)
}

func incorrectTypeError() {
	fmt.Println("Error: First parameter is 2 letters: either 'f' for floats, 'i' for ints, then 'u' for uniform and 'n' for normal")
	os.Exit(3)
}
