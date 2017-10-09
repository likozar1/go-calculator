package main

import "fmt"
import "regexp"
import "flag"
import "os"
import "bufio"
import "strconv"

const Regex = "-?[0-9]+|[*+/-]"

var Re *regexp.Regexp
var Expr string

func init() {
	flag.StringVar(&Expr, "ex", "", "Provide mathematical expression; example: '-10 + 15 * 3'")
	Re = regexp.MustCompile(Regex)
}

func main() {
	flag.Parse()

	if Expr == "" {
		chexpr := make(chan string, 1)

		defer close(chexpr)

		go readStdin(chexpr)

		Expr = <-chexpr
	}

	result := calculateParsedExpression(parseExpression(Expr))
	fmt.Println(Expr + " = " + result)
}

func readStdin(c chan string) {
	fmt.Print("Enter expression: ")
	b, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic(err)
	}

	c <- string(b)
}

func parseExpression(expression string) []string {
	parsed := Re.FindAllString(expression, -1)
	
	if len(parsed) % 2 != 1 {
		panic("invalid expression")
	}

	return parsed
}

func calculate(first int, second int, operator string) int {
	result := 0
	switch {
    case operator == "+":
        result = first + second
    case operator == "-":
        result = first - second
    case operator == "*":
		result = first * second
	case operator == "/":
		if second == 0 {
			panic("Division by zero")
		}
		result = first / second
	default:
		panic("undefined operator")
	}

	return result
}

func calculateParsedExpression(expressionArray []string) string {
	for i := 1; i < len(expressionArray); i=i+2 {
		first, err := strconv.Atoi(expressionArray[i-1])
		second, err := strconv.Atoi(expressionArray[i+1])

		if err != nil {
			panic("invalid expression")
		}

		expressionArray[i+1] = strconv.Itoa(calculate(first, second, expressionArray[i]))
		expressionArray[i-1] = ""
		expressionArray[i] = ""
	}

	return expressionArray[len(expressionArray)-1]
}
