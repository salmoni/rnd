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
	"time"
)

func main() {
	var dataFloat []float64
	var dataInt []int

	programArgs := os.Args
	lenArgs := len(programArgs)
	if lenArgs < 3 {
		// not enough parameters
		notEnoughParametersError()
	}
	seqNumber, error := strconv.Atoi(programArgs[3])
	if error != nil {
		incorrectNumberError()
	}

	seqDataType := programArgs[1]
	seqDataDist := programArgs[2]

	// Set random number generator seed to current date stamp
	t := time.Now().UnixNano() * 1000
	rand.Seed(t)

	if seqDataType == "f" {
		// generate a list of floats

		if lenArgs == 4 {
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
		} else if lenArgs == 6 {
			seqMin, error := strconv.ParseFloat(programArgs[4], 64)
			if error != nil {
				cannotConvertParameterError()
			}
			seqMax, error := strconv.ParseFloat(programArgs[5], 64)
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
					value := (rand.NormFloat64() * seqMin) + seqMax
					dataFloat = append(dataFloat, value)
				}
				printFloat(dataFloat)
			} else if seqDataDist == "e" {
				// exponentially distributed ints
				for idx := 0; idx < seqNumber; idx++ {
					value := (rand.ExpFloat64() * seqMin) + seqMax
					dataFloat = append(dataFloat, value)
				}
				printFloat(dataFloat)
			} else {
				// Error - didn't specify distribution
				incorrectDistributionError()
			}
		}
	} else {
		// Generate a list of integers

		if lenArgs == 6 {
			seqMin, error := strconv.ParseFloat(programArgs[4], 64)
			if error != nil {
				cannotConvertParameterError()
			}
			seqMax, error := strconv.ParseFloat(programArgs[5], 64)
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
				// Uniform distribution of ints
				for idx := 0; idx < seqNumber; idx++ {
					value := (rand.Float64() * seqRange) + seqMin
					dataInt = append(dataInt, int(value))
				}
				printInt(dataInt)

			} else if seqDataDist == "n" {
				// normal distribution of ints
				for idx := 0; idx < seqNumber; idx++ {
					value := (rand.NormFloat64() * seqMin) + seqMax
					dataInt = append(dataInt, int(value))
				}
				printInt(dataInt)
			} else if seqDataDist == "e" {
				// exponentially distributed ints
				for idx := 0; idx < seqNumber; idx++ {
					value := rand.ExpFloat64() / seqRange
					dataInt = append(dataInt, int(value))
				}
				printInt(dataInt)
			} else {
				// Error - didn't specify distribution
				incorrectDistributionError()
			}

		}
	}
}

func printInt(data []int) {
	// Print out ints
	n := len(data)
	if n < 1 {
		os.Exit(0)
	} else if n == 1 {
		fmt.Println(data[0])
	} else {
		outstr := strconv.Itoa(data[0])
		for idx := 1; idx < n; idx++ {
			outstr = outstr + ", " + strconv.Itoa(data[idx])
		}
		fmt.Println(outstr)
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

func showHelp() {
	fmt.Println("Missing parameters: type (float/int) distribution (uniform/normal) number [min/stdev] [max/mean]")
	os.Exit(3)
}

func notEnoughParametersError() {
	fmt.Println("Missing parameters: type (float/int) distribution (uniform/normal) number [min/stdev] [max/mean]")
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
