package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	
	"bytes"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

func getHashNo(strId string, nHashMax byte) byte {
    var bufs bytes.Buffer
     wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
     wr.Write([]byte(strId))
     wr.Close()
 
    convVal := bufs.String()
 
    nNo := byte(0)
    if len(convVal) > 0 {
        nNo = convVal[0]
        if nNo >= 'A' && nNo <= 'Z' {
            nNo = (nNo - 'A') + 'a'
        }
    }

	return nNo % nHashMax;
}


func getHashNo2(strId string, nHashSize byte) byte {   
	bufs, _, err := transform.String(korean.EUCKR.NewEncoder(), strId)
	if err != nil {
		 fmt.Println(err)
		 panic(err)        
	}
 
	 nNo := byte(bufs[0])
 
	 if nNo >= 'A' && nNo <= 'Z' {
		 nNo = (nNo - 'A') + 'a'
	 }
 
	 return nNo % nHashSize
}

func getHashPair(id string) (byte, byte) {
	return getHashNo(id, byte(6)), getHashNo2(id, byte(6))
}

func HashWrite(){
    var err error

	excelFileName := "./Data.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
        fmt.Printf(err.Error())
	}

	for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
			
			keyValue := row.Cells[0].String()

			hash1, hash2 := getHashPair(keyValue)

			if hash1 != hash2 {
				fmt.Printf("%s not Equal Hash %d %d\n", keyValue , hash1, hash2)
			}
			
			cellLen := len(row.Cells)
			if cellLen < 2 {
				row.AddCell()
			}

			if cellLen < 3 {
				row.AddCell()
			}

			row.Cells[1].SetInt(int(hash1))
			row.Cells[2].SetInt(int(hash2))
        }
	}
	
	err = xlFile.Save(excelFileName);
    if err != nil {
        fmt.Printf(err.Error())
    }
}

func main() {
	HashWrite()

	fmt.Printf("Work Complete")
}