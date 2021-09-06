package internal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type ArtHubTagData struct {
	ID       int64
	Version  string
	Topic    string
	Timeline string
	Type     string
	Style    string
	Color    string
	IsSync   bool
}

type ChildNodeRequestData struct {
	ParentID    int64             `json:"parent_id"`
	Offset      int64             `json:"offset"`
	Count       int64             `json:"count"`
	Filter      []string          `json:"filter"`
	Order       map[string]string `json:"order"`
	IsRecursive bool              `json:"is_recursive"`
}

type ChildNodeResp struct {
	Code   int64 `json:"code"`
	Result struct {
		Nodes []int64 `json:"nodes"`
		Order map[string]string
	} `json:"result"`
}

type AssetTagData struct {
	AssetID int64    `json:"asset_id"`
	TagName []string `json:"tag_name"`
}

type UploadResp struct {
	Code   int64   `json:"code"`
	Result []int64 `json:"result"`
}

func ReadExcel(path string, header map[string]string) error {
	fs, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}

	uploadResp := new(UploadResp)
	fc := excelize.NewFile()
	index := fc.NewSheet("Sheet1")

	rows, err := fs.GetRows("Sheet1")
	for i, row := range rows {
		log.Printf("index: %d, total: %d, data: %v", i+1, len(rows), row)
		if row[len(row)-1:][0] == "否" {
			row[len(row)-1] = "是"

			// 拼接request body
			FolderID, err := strconv.ParseInt(row[0], 10, 64)
			if err != nil {
				return err
			}

			childNodeRequestData := ChildNodeRequestData{
				ParentID:    FolderID,
				Offset:      0,
				Count:       -1,
				Filter:      []string{},
				Order:       map[string]string{"meta": "updated_date", "type": "descend"},
				IsRecursive: false,
			}

			childNodeRequestByteData, err := json.Marshal(childNodeRequestData)
			if err != nil {
				return err
			}

			// 请求子节点ID
			url := fmt.Sprintf("%s/%s/data/openapi/v2/core/get-child-node-id-in-range", GlobalConfig.Domain, GlobalConfig.Depot)
			resp, err := Post(url, bytes.NewReader(childNodeRequestByteData), header)
			if err != nil {
				return err
			}

			childNodeResp := new(ChildNodeResp)
			if err := json.Unmarshal(resp, childNodeResp); err != nil {
				return err
			}

			if childNodeResp.Code != 0 {
				return errors.New(fmt.Sprintf("get child nodes error, code is %d\n", childNodeResp.Code))
			}

			for _, v := range childNodeResp.Result.Nodes {
				assetData := AssetTagData{
					AssetID: v,
					TagName: row[3:7],
				}

				byteData, err := json.Marshal(assetData)
				if err != nil {
					return err
				}

				url := fmt.Sprintf("%s/%s/data/openapi/v2/core/add-asset-tag", GlobalConfig.Domain, GlobalConfig.Depot)
				resp, err := Post(url, bytes.NewReader(byteData), header)
				if err != nil {
					return err
				}

				if err := json.Unmarshal(resp, uploadResp); err != nil {
					return err
				}

				if uploadResp.Code != 0 {
					return errors.New(fmt.Sprintf("upload error, code is %d\n", uploadResp.Code))
				}
			}
		}

		if err := fc.SetSheetRow("Sheet1", fmt.Sprintf("A%d", i+1), &row); err != nil {
			return err
		}

		fc.SetActiveSheet(index)
		if err := fc.SaveAs("./new_" + GlobalConfig.File.Name); err != nil {
			return err
		}
	}
	return nil
}
