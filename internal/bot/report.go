package bot

import (
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Report map[string][]string

func (r Report) ToExcel() (io.Reader, error) {
	var (
		file      = excelize.NewFile()
		sheetName = "Sheet1"
		yIdx      = 1
		err       error
	)

	for username, links := range r {
		cell := "A" + strconv.Itoa(yIdx) // A1, A2, A3, etc.
		if err = file.SetCellStr(sheetName, cell, username); err != nil {
			return nil, err
		}

		yIdx++
		// links
		for _, link := range links {
			cell := "B" + strconv.Itoa(yIdx) // B1, B2, B3, etc.
			if err = file.SetCellStr(sheetName, cell, link); err != nil {
				return nil, err
			}
			yIdx++
		}
	}

	buf, err := file.WriteToBuffer()
	return buf, nil
}
