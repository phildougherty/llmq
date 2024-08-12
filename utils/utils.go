package utils

import (
    "bufio"
    "os"
    "strings"
)

func ReadStdin() (string, error) {
    scanner := bufio.NewScanner(os.Stdin)
    var input strings.Builder
    for scanner.Scan() {
        input.WriteString(scanner.Text() + "\n")
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

    return input.String(), nil
}
