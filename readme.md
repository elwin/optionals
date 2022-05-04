# Optionals in Go

Good idea? Not me to judge, we'll see.

## Instructions
```shell
go get github.com/elwin/optionals
```

## Example

```go
import . "github.com/elwin/optionals"

func someComputation(withResult bool) Optional[string] {
    if withResult {
        return Some("hello world")
    } else {
        return None[string]()
    }
}

func main() {
    someComputation(true).GetOrElse("default") // "hello-world"
    
    someComputation(false).GetOrElse("default") // "default"
    
    // or alternatively
    result := someComputation(true)
    if result.Exists() {
        result.Get() // "hello world"
    }
}
```