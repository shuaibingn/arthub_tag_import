package main

import (
	"log"

	"arthub_tag_import/internal"
)

func main() {
	header := map[string]string{"publictoken": internal.GlobalConfig.Token}
	if err := internal.ReadExcel(internal.GlobalConfig.File.Path+"\\"+internal.GlobalConfig.File.Name, header); err != nil {
		log.Println(err)
	}
}
