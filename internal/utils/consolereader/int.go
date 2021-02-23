package consolereader



import (
    "strconv"
    "strings"
)



// Function reads the input from a console and returns the value
func ReadConsoleInt64() (int64) {
    data, _ := r.ReadString('\n')
    data = strings.Replace(data, "\r\n", "", -1) // Note: Windows fix
    data = strings.Replace(data, "\n", "", -1) // Note: Windows fix
    dataInt, _ := strconv.ParseInt(data, 10, 64)
    return dataInt
}
