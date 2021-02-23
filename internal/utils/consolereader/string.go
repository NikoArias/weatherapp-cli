package consolereader



import (
    "strings"
)



// Function reads the input from a console and returns the value
func ReadConsoleString() (string) {
    data, _ := r.ReadString('\n')
    data = strings.Replace(data, "\r\n", "", -1) // Note: Windows fix
    return data
}
