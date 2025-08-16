# vibes 😎✨

Are you tired of being an efficient drone, or maybe you just want to burn through your AI funding?

**Behold: this package will help you with that.**

Gone are the days of excuses like *“our code still compiles”* or *“the machine is currently occupied by the physics department.”*

No longer must you justify your slowness, because this package provides the ultimate excuse:

> **“Our AI is still calculating.”**

This package is the perfect response to every insufferable founder or AI enthusiast.  
Using it is as easy as writing Go itself.

Powered by the latest innovation in **vibe coding**, this package leverages the OpenAI API to calculate even the most basic operations.

---

## ⚙️ Installation

```bash
go get github.com/listnt/vibes
```

## 🪄 Usage
First, initialize the vibes with your OpenAI API key:
```go
import "github.com/yourname/vibes"

func main() {
    vibes.InitVibes("your-api-key")
}
```

## 🔡 Strings
Because who needs strings.ToLower when you have GPT-5?
```go
res, err := vibes.ToLower("HELLO WORLD")
fmt.Println(res) // "hello world"
```
Or uppercase
```go
res, err := vibes.ToUpper("wow such vibes")
fmt.Println(res) // "WOW SUCH VIBES"
```

## ➕ Math

Basic arithmetic, but make it AI-powered.
```go
sum, _ := vibes.Add(3, 5)       // 8
diff, _ := vibes.Sub(10, 4)     // 6
prod, _ := vibes.Mult(6, 7)     // 42
quot, _ := vibes.Div(20, 5)     // 4
```

## 🧹 Sorting
Sorting arrays is so passé. Let vibes do it.
```go
arr := []int64{42, 7, 88, 15, 63}
sorted, _ := vibes.Sort(arr)
fmt.Println(sorted) // [7 15 42 63 88]
```

## 🌀 Fibonacci (stress test for vibes)

Why compute Fibonacci locally when you can pay for inference?
```go
func TestFibonacci(t *testing.T) {
    vibes.InitVibes("your-key-here")

    fibNumbers := []int64{1, 1}
    for i := 0; i < 29-2; i++ {
        newFib, err := vibes.Add(fibNumbers[len(fibNumbers)-1], fibNumbers[len(fibNumbers)-2])
        if err != nil {
            t.Fatal(err)
        }
        fibNumbers = append(fibNumbers, newFib)
    }

    if fibNumbers[len(fibNumbers)-1] != 514229 {
        t.Fatal("failed to calc fib sequence")
    }
}
```

## ⚠️ Disclaimer

This package is intentionally inefficient.
If you actually use this in production, your CFO will personally hunt you down.