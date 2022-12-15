package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func Scaffold(problemNumber string) {

	// Create the new directory
	solveDir := filepath.Join("solves", problemNumber)
	err := os.Mkdir(solveDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory %s: %s", solveDir, err)
		return
	}

	// Create the new file
	solveFile := filepath.Join(solveDir, "solve.go")
	f, err := os.Create(solveFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %s", solveFile, err)
		return
	}

	// Write the file contents
	_, err = f.WriteString(fmt.Sprintf(`package solve%s

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	dirname := filepath.Dir(filename)

	dat, err := os.ReadFile(filepath.Join(dirname, "../../inputs/%s.txt"))
	if err != nil {
		panic(err)
	}

	input := string(dat)

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input string) int {
	return 0
}

func PartTwo(input string) int {
	return 0
}
`, problemNumber, problemNumber))

	if err != nil {
		fmt.Printf("Error writing to file %s: %s", solveFile, err)
		return
	}

	// Close the file
	err = f.Close()
	if err != nil {
		fmt.Printf("Error closing file %s: %s", solveFile, err)
		return
	}

	// Create the test file
	testFile := filepath.Join(solveDir, "solve_test.go")
	f, err = os.Create(testFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %s", testFile, err)
		return
	}

	template, _ := template.New("testFileContent").Parse(`
package solve{{ .Number }}_test

import (
	"io/ioutil"
	"testing"

	solve{{ .Number }} "github.com/plastic041/aoc-2022-go/solves/{{ .Number }}"
	"github.com/stretchr/testify/assert"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestPartOne(t *testing.T) {
	dat, err := ioutil.ReadFile("../../examples/{{ .Number }}.txt")
	check(err)

	assert.Equal(t, 0, solve{{ .Number }}.PartOne(string(dat)))
}

func TestPartTwo(t *testing.T) {
	dat, err := ioutil.ReadFile("../../examples/{{ .Number }}.txt")
	check(err)

	assert.Equal(t, 0, solve{{ .Number }}.PartTwo(string(dat)))
}
`)

	testTextContent := struct {
		Number string
	}{
		Number: problemNumber,
	}

	err = template.Execute(f, testTextContent)

	if err != nil {
		fmt.Printf("Error writing to file %s: %s", testFile, err)
		return
	}

	// Close the file
	err = f.Close()
	if err != nil {
		fmt.Printf("Error closing file %s: %s", testFile, err)
		return
	}

	// Create the example file
	exampleFile := filepath.Join("examples", problemNumber+".txt")
	f, err = os.Create(exampleFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %s", exampleFile, err)
		return
	}

	// Close the file
	err = f.Close()
	if err != nil {
		fmt.Printf("Error closing file %s: %s", exampleFile, err)
		return
	}

	// Create the input file
	inputFile := filepath.Join("inputs", problemNumber+".txt")
	f, err = os.Create(inputFile)
	if err != nil {
		fmt.Printf("Error creating file %s: %s", inputFile, err)
		return
	}

	// Close the file
	err = f.Close()
	if err != nil {
		fmt.Printf("Error closing file %s: %s", inputFile, err)
		return
	}

	fmt.Printf("Scaffolded problem %s", problemNumber)
}
