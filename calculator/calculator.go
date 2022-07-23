package main

import (
	"errors"
	"fmt"
	mathFunction "github.com/kaldiroglu/Fundamentals-of-Go/ch07/math"
	"math"
	"os"
	"strconv"
	"strings"
)

var flag bool = true

type Calculator struct {
	functions []mathFunction.MathFunction
}

func (c *Calculator) addMathFunction(m mathFunction.MathFunction) {
	c.functions = append(c.functions, m)
}

func (c *Calculator) doCalculation(name string, arg float64) (float64, error) {
	var result float64
	for _, f := range c.functions {
		if strings.ToUpper(name) == strings.ToUpper(f.GetName()) {
			result = f.Calculate(arg)
			return result, nil
		}
	}
	return 0, errors.New("invalid number argument:" + strconv.FormatFloat(arg, 'E', -1, 64))
}

func (c *Calculator) checkValidFunction(name string) error {
	for _, f := range c.functions {
		if strings.ToUpper(name) == strings.ToUpper("x") {
			fmt.Println("Bye!")
			os.Exit(0)
		} else if strings.ToUpper(name) == strings.ToUpper(f.GetName()) {
			return nil
		}
	}
	return errors.New("invalid function name:" + name)
}

func main() {
	//functions()
	//calculator()
	startCalculator()
}

func functions() {
	sin := mathFunction.Sin{"Sinus"}
	fmt.Printf("%v\n", sin)
	sin30 := sin.Calculate(math.Pi / 6)
	fmt.Printf("Sinus of 30 degree is %f\n", sin30)

	cos := mathFunction.Cos{"Cosinus"}
	fmt.Printf("%v\n", sin)
	cos30 := cos.Calculate(math.Pi / 6)
	fmt.Printf("Cosinus of 30 degree is %f\n", cos30)

	log := mathFunction.Log{"Log of base e"}
	fmt.Printf("%v\n", log)
	logE := log.Calculate(2.71828)
	fmt.Printf("Log of Euler constant is %f\n", logE)

	var mf1 mathFunction.MathFunction = sin
	fmt.Printf("%v\n", mf1)

	mf1 = cos
	fmt.Printf("%v\n", mf1)

	mf1 = log
	fmt.Printf("%v\n", mf1)
}

func calculator() {
	myCalculator := Calculator{}
	//myCalculator.functions[0] = Sin{"Sinus"}
	//myCalculator.functions[1] = Cos{"Cosinus"}
	//myCalculator.functions[2] = Log{"Log"}

	myCalculator.addMathFunction(mathFunction.Sin{"Sinus"})
	myCalculator.addMathFunction(mathFunction.Cos{"Cosines"})
	myCalculator.addMathFunction(mathFunction.Log{"Log"})

	fmt.Println(myCalculator.doCalculation("Sinus", math.Pi/6))
	fmt.Println(myCalculator.doCalculation("Cosines", math.Pi/6))
	fmt.Println(myCalculator.doCalculation("Log", math.E))
}

func startCalculator() {
	myCalculator := Calculator{}
	//myCalculator.functions[0] = Sin{"Sinus"}
	//myCalculator.functions[1] = Cos{"Cosinus"}
	//myCalculator.functions[2] = Log{"Log"}
	myCalculator.addMathFunction(mathFunction.Sin{"Sine"})
	myCalculator.addMathFunction(mathFunction.Cos{"Cosine"})
	myCalculator.addMathFunction(mathFunction.Log{"Log"})

	fmt.Println("\nCalculator started.")
	fmt.Println("You can calculate using following functions")
	for _, f := range myCalculator.functions {
		fmt.Println(f.GetName())
	}

	for flag {
		//var x string
		//_, err := fmt.Scanf("%s", &x)
		//if err != nil {
		//	fmt.Println(err)
		//	os.Exit(0)
		//}
		var fName string
		var arg float64
		fmt.Println("> Enter name of the calculation or enter x to exit:")
		_, err := fmt.Scanln(&fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		err = myCalculator.checkValidFunction(fName)
		if err != nil {
			fmt.Println(err)
			fmt.Println()
			continue
		}
		//
		if fName == "x" {
			//os.Exit(0)
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanln(&arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			value, err := myCalculator.doCalculation(fName, arg)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Result of %s of %f : %f\n", fName, arg, value)
				fmt.Println()
			}
		}
	}
	println("Bye!")
}
