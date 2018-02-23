package main

/*
#include "stdio.h"
#include <string>
using namespace std;

int getHashNo(const string& strId, int nHashSize)
{
	unsigned char nNo = (unsigned char)(*strId.c_str());
	if (nNo >= 65 && nNo <= 90)
	{
		nNo = (nNo - 65) + 97;
	}

	return nNo % nHashSize;
}

*/
import "C"
import "fmt"

func main() {
	hashNo := C.getHashNo(C.CString("nupnS02"), C.int(6))
	fmt.Println(hashNo);
}

