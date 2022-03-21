// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
)

func GetAPIKey(key string) string {
	data, err := ioutil.ReadFile("keys.json")
	if err != nil {
		panic(err)
	}
	var returnData map[string]interface{}
	err = json.Unmarshal(data, &returnData)
	if err != nil {
		fmt.Printf("%+v", err.Error())
		return ""
	}
	return returnData[key].(string)
}

func Runner(email string, kind string) {
	cmd := exec.Command("python3", "modules/"+kind+".py", "-e", email)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)
	cmd.Wait()

}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
