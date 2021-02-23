package consolereader



import (
    "os"
    "bufio"
)



var r *bufio.Reader



func init() {
    r = bufio.NewReader(os.Stdin)
}
