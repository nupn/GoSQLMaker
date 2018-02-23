package main



import (
 "bytes"
 "fmt"
 "golang.org/x/text/encoding/korean"
 "golang.org/x/text/transform"
)


func GetHashNo(strId string, nHashMax uint) uint {
    var bufs bytes.Buffer
     wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())
     wr.Write([]byte(strId))
     wr.Close()
 
    convVal := bufs.String()
 
    var nNo uint
    nNo = 0
    if len(convVal) > 0 {
        nNo = uint(convVal[0])
        if nNo >= 'A' && nNo <= 'Z' {
            nNo = (nNo - 'A') + 'a';
        }
    }

	return nNo % nHashMax;
}

//참고
//https://godoc.org/golang.org/x/text/encoding
//http://blog.naver.com/nersion/220884742148

//Test

func main() {
    nHash := GetHashNo("nupns2", uint(6))
    fmt.Println(nHash);


    /*
    testString := "크아"

    for i := 0; i < len(testString); i++ {

        printf("%d", testString[i])

    }

    printf("\n")

    printf("After Convert EUC-KR\n")



    var bufs bytes.Buffer
   // transform.String(korean.EUCKR.NewEncoder())
    wr := transform.NewWriter(&bufs, korean.EUCKR.NewEncoder())

    wr.Write([]byte(testString))

    wr.Close()

    convVal := bufs.String()

    for i := 0; i < len(convVal); i++ {

        printf("%d ", convVal[i])

     }

     printf("\n")
     */

}