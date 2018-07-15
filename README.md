# rnd
A command line program in Go to produce a list of random numbers

## Usage

rnd [parameters] type distribution number min/stddev max/mean

type = type of data: 'f' for float (64 bit), 'i' for integers
distribution = 'u' = uniform distribution, 'n' = normal distribution
number = how many numbers to produce (must be int)
min/stddev = either minimum value (for uniform distributions) or standard deviation (for normal distributions)
max/mean = either maximum value (for uniform distributions) or mean (for normal distributions)

## Copyright information

(c) 2018 Alan James Salmoni. Released under the MIT licence
