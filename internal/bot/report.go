package bot

import (
	"io"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Report map[string][]string

func (r Report) ToExcel() (io.Reader, error) {
	var (
		file               = excelize.NewFile()
		sheetName          = "Sheet1"
		yIdx               = 1
		err                error
		longestUsernameLen = 0
		longestLinkLen     = 0
		styleId            = 0
	)

	if styleId, err = file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Pattern: 1,
			Color:   []string{"#FFFF00"},
		},
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	}); err != nil {
		return nil, err
	}

	for username, links := range r {
		cell := "A" + strconv.Itoa(yIdx) // A1, A2, A3, etc.
		if err = file.SetCellStr(sheetName, cell, username); err != nil {
			return nil, err
		}
		if err = file.SetCellStyle(sheetName, cell, cell, styleId); err != nil {
			return nil, err
		}

		longestUsernameLen = max(longestUsernameLen, len(username))
		yIdx++
		// links
		for _, link := range links {
			cell := "B" + strconv.Itoa(yIdx) // B1, B2, B3, etc.
			if err = file.SetCellStr(sheetName, cell, link); err != nil {
				return nil, err
			}

			longestLinkLen = max(longestLinkLen, len(link))
			yIdx++
		}
	}

	if err = file.SetColWidth(
		sheetName, "A", "A",
		float64(longestUsernameLen+2), // +2 for margin
	); err != nil {
		return nil, err
	}

	if err = file.SetColWidth(
		sheetName, "B", "B",
		float64(longestLinkLen+2), // +2 for margin
	); err != nil {
		return nil, err
	}

	buf, err := file.WriteToBuffer()
	return buf, nil
}
