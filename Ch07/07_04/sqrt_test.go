// Challenge : Test your SQRT function with values from a CSV file. The CSV file
// will have a value and expected result in every line.
package sqrt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

type testCase struct {
	value	float64
	expected	float64
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)

			if err != nil {
				t.Fatalf("error in calculation - %s", err)
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected);
			}
		})
	}
}

func TestFromFile(t *testing.T) {

	// [START : Parse the test case configuration file]
	var testCasesFilePath = "sqrt_cases.csv";
	file, err := os.Open(testCasesFilePath);
	if err != nil {
		t.Fatalf("Cannot open test cases file %s : %s", testCasesFilePath, err);
	}

	defer file.Close(); // close the file descriptor

	var testCases = make([]testCase,0,10); // allocate slice of test cases
	var scanner = bufio.NewScanner(file); // we want to scan the file line by line

	csvSplitFunc := func(c rune) bool { // and split de content of the line betweeen comma
		return c == ',';
	}

	for lineIndex :=0; scanner.Scan(); lineIndex++ {
		tokens := strings.FieldsFunc(scanner.Text(), csvSplitFunc);
		if len(tokens) != 2 { // we expect only 2 values
			t.Fatalf("test cases file format expects 2 values by line. review file <%s> Line <%d>", testCasesFilePath, lineIndex);
		}

		tcValue, err := strconv.ParseFloat(tokens[0], 64); //convert string value to Float64
		if err != nil {
			t.Fatalf("test case Value on line <%d> : <%s> cannot be converted to Float64", lineIndex, tokens[0]);
		}
		tcExpectedValue, err := strconv.ParseFloat(tokens[1], 64);
		if err != nil {
			t.Fatalf("test case Expected Value on line <%d> : <%s> cannot be converted to Float64", lineIndex, tokens[1]);
		}

		// fill in the testCases slice
		testCases = append(testCases, testCase{
			value: tcValue,
			expected: tcExpectedValue,
		})

	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("test cases scan error : <%s>", err)
	}

	// [END : Parse the test case configuration file]

	// [START : Run test cases]
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)

			if err != nil {
				t.Fatalf("error in calculation - %s", err)
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected);
			}
		})
	}
}
