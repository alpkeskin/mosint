// mosint v2.1
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Website: https://imalp.co
package modules

import (
	"log"
	"os"
	"strconv"
	"time"
)

func FileWriter(email string, text string) string {
	t := time.Now()

	var filename string = "outputs/" + email + "-" + strconv.FormatUint(uint64(t.Unix()), 10) + ".txt"
	f, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString(text)
	return filename
}
