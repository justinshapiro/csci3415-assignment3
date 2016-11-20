# Program 3: Go
##### Problem Statement
For the third programming assignment, the task was to create a simple calculator in the Go programming language. By “simple”, we mean that the calculator only should support addition, subtraction, multiplication and division operations with both floating-point and integer numbers. In addition, the calculator must also support parenthesis. Any characters taken as input besides +, -, *, /, (, ), and 0-9 will throw an error. The precedence rules for this calculator are that multiplication and division are of higher precedence than addition and subtraction, but the parenthesis will always override this precedence by evaluating inside any existing parentheses first.  Evaluation of expressions with both integer and floating-point numbers will be done using reflection .

##### Implementation
A large part of this implementation uses Dr. Williams’ code that he provided. The code provided specified a command-line calculator that satisfies most of the problem statement, only leaving out floating point calculations and any associated use of reflection, parentheses, and some error checking. My implementation takes this existing code and modifies it to support the following functionality:
1.	Mixed floating point and integer calculations with the help of reflection.
    *	The compute function was changed to return an interface{} to the main function support this.
    *	The same operand stack was used to store both integer and floating-point values, rather than having a separate operand stack for floating point values.
    *	Specifically, int64 and float64 are the base data types used for integer and floating-point values, respectively. These types are checked during runtime by using reflection package functions reflect.Kind(), to check for reflect.Int64 and reflect.Float64, and reflect.valueOf() to make the int64 or float64 value a reflection type.
2.	The use of parentheses in expressions that can have any number of sets of parentheses embedded inside them.
    *	Done using recursion. A new string is created within the compute function with the expression inside the given set of parenthesis. This string is a parameter to a recursive call to the compute function. Since the compute function will return the result of this expression, when the recursive call finishes the value of the expression inside the parenthesis will be placed in the original operand stack.
    *	To support recursion, operatorStack and operandStack were moved to be inside (local to) the compute function so that a recursive call to compute will make it so the evaluation of the inside of a given set of parenthesis will have its own set of stacks.
    *	As another requirement to support recursion, the apply function was modified to return a copy of the modified pair of stacks for a given recursive call. This makes it so compute can return the value of an expression.
3.	The error checking was implemented:
    *	Disallows non-operator characters in front of the start of a set of parentheses.
    *	Disallows spaces in between two operands.
    *	For an illegal character error, the error message will have that character in it to let the user know where the syntax error was.