package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"arthub_tag_import/internal"
)

func main() {
	header := map[string]string{"publictoken": internal.GlobalConfig.Token}
	path := strings.ReplaceAll(internal.GlobalConfig.File.Path, "\\", "/")
	if err := internal.ReadExcel(path+"/"+internal.GlobalConfig.File.Name, header); err != nil {
		log.Println(err)
	}
	log.Println("finished!!!")
	log.Println("press any key to continue...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}
