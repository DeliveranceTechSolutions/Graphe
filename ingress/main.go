package ingress

import (
	"io/fs"
)

type Reader struct {
	buffer []byte
}

func feedCSVData(csv fs.File) {
	
	csv.Read([]byte{})
} 