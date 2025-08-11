# Vibesort

**Vibesort** is a fun Go package that uses OpenAI's GPT-5 model to sort integers — not through conventional algorithms, but by *asking the AI nicely* to put them in order.

Think of it as bubble sort with vibes.

---

## ✨ Features
- Sorts integers in ascending order using GPT-5.
- Simple API: pass in a slice of `int64` and get a sorted slice back.
- Demonstrates how to integrate OpenAI's API into a Go project.

---

## 📦 Installation

Make sure you have Go installed, then:

```bash
go get github.com/openai/openai-go/v2
```
Clone or copy this package into your project.

## 🔑 Setup
You’ll need an OpenAI API key.
Get one from: https://platform.openai.com/

Set it as an environment variable:

```bash
export OPENAI_API_KEY="your-secret-key"
```

## 🛠 Usage
Example:

```go 
package main

import (
	"fmt"
	"log"
)

func main() {
	v := NewVibesort("your-api-key")

	numbers := []int64{42, 7, 88, 15, 63}

	sorted, err := v.vibesort(numbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sorted) // [7 15 42 63 88]
}
```

## 🧪 Testing
The included test (vibesort_test.go) checks that the sorting works as expected.

Run:

```bash
go test -v
```
Example output:

```
=== RUN   TestInts
[7 15 42 63 88] <nil>
--- PASS: TestInts (0.45s)
PASS
ok  	vibesort	0.48s
```

## ⚙️ How It Works
The package sends your unsorted list of integers to GPT-5.

GPT-5 responds with a sorted list in the format [1 2 3 ...].

The package parses this response back into []int64.

Note: This is not the fastest sorting algorithm — network calls are involved — but it is the most polite.

## ⚠️ Caveats
Requires internet access and a valid OpenAI API key.

Response parsing assumes GPT strictly follows the [1 2 3 ...] format.

API usage may incur costs depending on your OpenAI plan.

