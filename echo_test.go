package echo

import (
    "encoding/hex"
    "fmt"
    "testing"
    "os"
    "strings"
    "strconv"
    "bufio"
    "hash"
    "bytes"
    "io"
)

func testFromKAT(t *testing.T, filename string, bit_length int) {
    f, err := os.Open(filename)
    if err != nil {
        t.Error(fmt.Sprint("Unable to open test file", filename))
        return
    }
    defer f.Close()

    reader := bufio.NewReader(f)
    var (
        length int
        msg []byte
        hash_expected []byte
    )
    for {
        part, _, err := reader.ReadLine()

        if err != nil {
            if err != io.EOF {
                t.Error("error")
            }
            break
        }
        line := string(part)
        var h hash.Hash
        switch {
        case strings.Index(line, "Len = ") == 0:
            length, _ = strconv.Atoi(line[6:])
        case strings.Index(line, "Msg = ") == 0:
            msg, _ = hex.DecodeString(line[6:])
            if length/8 < len(msg) {
                msg = msg[:length/8]
            }
        case strings.Index(line, "MD = ") == 0:
            if length % 8 != 0 || len(msg) != length/8 {
                continue
            }
            hash_expected, _ = hex.DecodeString(line[5:])
            if len(hash_expected) != bit_length/8 {
                t.Error("error in sample file")
            }
            switch bit_length {
            case 224:
                h = New224()
            case 256:
                h = New256()
            case 384:
                h = New384()
            case 512:
                h = New512()
            }
            h.Write(msg)
            hash_actual := h.Sum(nil)
            if ! bytes.Equal(hash_actual, hash_expected) {
                t.Error("hashfail on actual=",  hex.EncodeToString(hash_actual), " expected=", hex.EncodeToString(hash_expected), "msg=", msg)
            }
        }
    }
}


func Test224(t *testing.T) {
    testFromKAT(t, "ShortMsgKAT_224.txt", 224)
    testFromKAT(t, "LongMsgKAT_224.txt", 224)
}

func Test256(t *testing.T) {
    testFromKAT(t, "ShortMsgKAT_256.txt", 256)
    testFromKAT(t, "LongMsgKAT_256.txt", 256)
}

func Test384(t *testing.T) {
    testFromKAT(t, "ShortMsgKAT_384.txt", 384)
    testFromKAT(t, "LongMsgKAT_384.txt", 384)
}

func Test512(t *testing.T) {
    testFromKAT(t, "ShortMsgKAT_512.txt", 512)
    testFromKAT(t, "LongMsgKAT_512.txt", 512)
}

// benchmarks

var x = func() []byte {
    result := make([]byte, 1024*1024)
    for i := 0; i < 1024*1024; i++ {
        result = append(result, 0)
    }
    return result
}()

func Benchmark1MiB224(b *testing.B) {
    h := New224()
    for i := 0; i < b.N; i++ {
        h.Write(x)
        h.Sum(nil)
    }
}

func Benchmark1MiB256(b *testing.B) {
    h := New256()
    for i := 0; i < b.N; i++ {
        h.Write(x)
        h.Sum(nil)
    }
}

func Benchmark1MiB384(b *testing.B) {
    h := New384()
    for i := 0; i < b.N; i++ {
        h.Write(x)
        h.Sum(nil)
    }
}

func Benchmark1MiB512(b *testing.B) {
    h := New512()
    for i := 0; i < b.N; i++ {
        h.Write(x)
        h.Sum(nil)
    }
}
