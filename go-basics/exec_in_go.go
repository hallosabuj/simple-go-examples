package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func FirstExample() {
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("and old falcon")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String())
}

func MultipleArgs() {
	//=================================================
	//exec command with multiple args
	var out bytes.Buffer
	cmd := exec.Command("echo", "Welcome", "to", "Golang")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %q\n", out.String())
}

func CaptureOutput() {
	//Capture output
	cmd := exec.Command("ls")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func ExampleCmd_Output() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

func ExampleLookPath() {
	var command = "ls"
	//Works as which command
	path, err := exec.LookPath(command)
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("%s is available at %s\n", command, path)
}

func ExampleCmd_Run() {
	cmd := exec.Command("sleep", "2")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func ExampleCmd_Start() {
	cmd := exec.Command("sleep", "2")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func ExampleCmd_StdinPipe() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}

func test() {
	c1 := exec.Command("echo", "allowme1")
	c2 := exec.Command("sudo", "-S", "apt", "update")

	var stderr bytes.Buffer
	pr, pw := io.Pipe()
	c1.Stdout = pw
	c2.Stdin = pr
	c2.Stdout = os.Stdout
	c2.Stderr = &stderr
	c1.Start()

	go func() {
		defer pw.Close()
		c1.Wait()
	}()
	fmt.Println()
	err := c2.Run()
	if err != nil {
		fmt.Println(stderr)
	}
}

func test1() {
	//Capture output
	cmd := exec.Command("echo", "Hello Wolrd", "|", "grep", "Hello")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func main() {
	test1()
}
