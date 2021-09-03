package test

import (
	"fmt"
	"testing"

	"arthub_tag_import/internal"
)

func TestExcelRead(t *testing.T) {
	header := map[string]string{"publictoken": "200e5"}

	err := internal.ReadExcel("E:\\go\\src\\arthub_tag_import\\arthub_test.xlsx", "http://arthub-api.h3d.com.cn/h3d/data/openapi/v2/core/add-asset-tag", header)
	fmt.Println(err)
}
