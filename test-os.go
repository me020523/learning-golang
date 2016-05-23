package main

import (
	"fmt"
	"os"
)

func getEnvGroup() {
	envs := os.Environ()
	for _, env := range envs {
		fmt.Println(env)
	}
}
func testExpand(name string) {
	testStr := fmt.Sprintf("hello, ${%s}", name)
	newStr := os.Expand(testStr, func(s string) string {
		return "My " + s
	})
	fmt.Println(newStr)
}
func getWorkDir() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
}
func getHostName() {
	hostname, _ := os.Hostname()
	fmt.Println(hostname)
}
func getEnv(name string) {
	value, ok := os.LookupEnv(name)
	if !ok {
		fmt.Printf("No Such Env Key: %s \n", value)
	}
	fmt.Println(value)
}
func linkTarget(name string) {
	target, _ := os.Readlink(name)
	fmt.Println(target)
}
func startProcess(command string, args []string) {
	atrr := &os.ProcAttr{
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	process, _ := os.StartProcess(command, args, atrr)
	/*if err != nil {
		fmt.Println(err.Error())
		return
	}*/
	process.Wait()
}
func main() {
	//testExpand("Time")
	//getWorkDir()
	//getHostName()
	//getEnv("NoPath")
	//getEnv("PATH")
	//getEnv("path")
	//getEnv("pAth")
	//linkTarget("test.link")
	startProcess("/bin/cat",
		[]string{
			"/bin/cat",
			"/Users/bcshuai/Documents/git/go/src/github.com/me020523/learning-golang/op-demo.go",
		})
}
