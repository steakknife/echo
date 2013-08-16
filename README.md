# ECHO hash in Go

### Why?

ECHO has very nice properties, it is slower that other SHA-3 round2/3 finalists in both hardware and software.

If you care about the function of cost for bruteforcing, you understand why this is important.

### Usage

    go get github.com/steakknife/echo


    
```go
import "github.com/steakknife/echo"
import "github.com/steakknife/securecompare"


h := echo.New256()
h.Write(someBytes)
result := h.Sum(nil)

if securecompare.Equal(result, expected) {
    // good
} else {
    // bad
}
```


### Acknowledgements

France telecom, you rock

echo.c and echo.h are Optimized_64 from http://csrc.nist.gov/groups/ST/hash/sha-3/Round2/documents/ECHO_Round2.zip 


### License

MIT

## Patches welcome!
