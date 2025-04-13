<p align="center">
  <a href="https://go.dev/" target="blank"><img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="120" alt="Go Logo" /></a>
</p>

<p align="center"><strong>A lightweight and efficient Golang project documenting the journey through core language fundamentals.</strong></p>

<p align="center">
<a href="https://go.dev/" target="_blank"><img src="https://img.shields.io/badge/Language-Go-blue.svg" alt="Go Language" /></a>
<a href="https://github.com/" target="_blank"><img src="https://img.shields.io/github/license/yourusername/your-repo.svg" alt="License" /></a>
<a href="https://github.com/" target="_blank"><img src="https://img.shields.io/github/stars/yourusername/your-repo?style=social" alt="GitHub stars" /></a>
<a href="https://twitter.com/golang" target="_blank"><img src="https://img.shields.io/twitter/follow/golang.svg?style=social&label=Follow" alt="Follow Go on Twitter"></a>
</p>


## Stage 1: Core Basics

These fundamental concepts lay the groundwork for writing efficient and readable Go code.

---

### Core Topics Covered

| Topic                    | Description                                                                 |
|--------------------------|-----------------------------------------------------------------------------|
|  Hello World            | My first Go program! Printed text to the console.                          |
|  Values                 | Learned about Go literals like strings, integers, booleans, etc.           |
|  Variables              | Declared and used variables using `var`, `:=`, and explored scoping.       |
|  Constants              | Defined constant values using `const`, including math with arbitrary precision. |
|  For                   | Wrote loops with `for`, the only looping construct in Go.                  |
|  If/Else                | Conditional logic with `if`, `else if`, and `else`.                        |
|  Switch                 | Used `switch` statements for cleaner branching logic.                      |
|  Arrays                 | Worked with fixed-length sequences of elements.                           |
|  Slices                 | Mastered flexible, dynamically-sized arrays using slices.                 |
|  Maps                   | Stored key-value pairs with maps and checked for key existence.            |
|  Functions              | Defined reusable blocks of code using `func`.                              |
|  Multiple Return Values | Used functions that return more than one value (including `ok` patterns).  |
|  Variadic Functions     | Created functions that accept a variable number of arguments.              |

---

### 🛠 Example Snippet

```go
// Variadic function example
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4)) // Output: 10
}
