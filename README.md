# week-4-homework-1


- Please read "A Guide to the Go Garbage Collector" at https://go.dev/doc/gc-guide and come up with questions you may have in the next class.
      
- Add error handling to the example calculator.go of ch07 we studied at the end of last class. And also make sure that the example works file even though math function name entered by the user is lower case of name registered by the calculator. i.e. make the calculator match math function names case insensitive. (In fact the new copy that I uploaded yesterday to Google Drive includes these new features bu you can try to implement yours and then have a look at the way I implemented 🙂)
      
- Create an abstraction to make conversions among different temperature systems such as Celsius, Fahrenheit and Kelvin. Create a calculator that is loosely coupled to that abstraction i.e. program to an interface, not an implementation principles is observed. Make sure that the program can interact with the user so that user can achieve some conversions using the program.
      
- How would you implement a program to list go source code files in a given directory including all sub-directories recursively? Please use filepath.WalkDir() function (https://pkg.go.dev/path/filepath#WalkDir) for this purpose.

## Do not forget to commit the initial and final work.
