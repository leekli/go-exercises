// Spawning Processes
// - Sometimes our Go programs need to spawn other, non-Go processes.
// - Using `os/exec` to run a process

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	// Spawning a process with a flag (omit 2nd argument if no flags needed)
	// exec the date command, save external process into `dateCmd object`
	dateCmd := exec.Command("date", "-jRnu")

	// Get the `dateCmd` output
	dateOut, err := dateCmd.Output()

    if err != nil {
        panic(err)
    }

    fmt.Println("> date", string(dateOut))

	// Spawning a process with a flag (error)
	// - Output and other methods of Command will return *exec.Error if there was a problem executing the command (e.g. wrong path), and *exec.ExitError if the command ran but exited with a non-zero return code.
	_, err = exec.Command("date", "-x").Output()

    if err != nil {
        switch e := err.(type) {
        case *exec.Error:
            fmt.Println("failed executing:", err)
        case *exec.ExitError:
            fmt.Println("command exit rc =", e.ExitCode())
        default:
            panic(err)
        }
    }

	// Another example using `ls`
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")

    lsOut, err := lsCmd.Output()

    if err != nil {
        panic(err)
    }

    fmt.Println("> ls -a -l -h", string(lsOut))

	// Using exec with piped data (`grep`)
	grepCmd := exec.Command("grep", "hello")

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := io.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))
}