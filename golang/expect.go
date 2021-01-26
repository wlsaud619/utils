package golang

import (
        "fmt"

        "log"
        "regexp"
        "time"
        "strings"

        "github.com/google/goexpect"
)

var (
        timeout = 10 * time.Second
        promptRE = regexp.MustCompile("%")
)


func ExpectSingle(cmd string, question string, answer string) string {

        e, _, err := expect.Spawn(fmt.Sprintf(cmd), -1)
        if err != nil {
                log.Fatal(err)
        }
        defer e.Close()

        e.Expect(regexp.MustCompile(question), timeout)
        e.Send(answer+ "\n")

        result, _, _ := e.Expect(promptRE, timeout)
        e.Send("exit\n")

        return strings.TrimSpace(result)
}

