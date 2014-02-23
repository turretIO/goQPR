package qpr 

import (
    "fmt"
    "bytes"
)

const (
    escape string = "="
    maxlinesize int = 76
    emptystring string = ""
    maxchar = "~"
    minchar = " "
)

func NewQPEncoder() (*QPEncoder) {
	qp := new(QPEncoder)
	return qp
}

type QPEncoder struct {
	counter		int
}

func (qp *QPEncoder) quote(b byte, force bool) ([]byte, int) {
	if(force) {
        return []byte(fmt.Sprintf("=%X", b)), len([]byte(fmt.Sprintf("=%X", b)))		
	}
	if b < []byte(minchar)[0] || b > []byte(maxchar)[0] {
        return []byte(fmt.Sprintf("=%X", b)), len([]byte(fmt.Sprintf("=%X", b)))
    }
    if b == []byte("=")[0] {
        return []byte(fmt.Sprintf("=%X", b)), len([]byte(fmt.Sprintf("=%X", b)))    	
    }

    return []byte(string(b)), len([]byte(string(b))) 
}


func (qp *QPEncoder) encodeLine(encoded *bytes.Buffer, line *[]byte) {
	*line = bytes.Replace(*line, []byte("\n"), []byte("\r\n"), -1)
	var buf bytes.Buffer
	for index, chr := range *line {
		enc, encLen := qp.quote(chr, false)
		if index == len(*line)-1 && chr == []byte(" ")[0]{
			enc, encLen = qp.quote(chr, true)			
		}
		qp.counter += encLen
		if qp.counter > maxlinesize-1 {
			buf.Write([]byte("=\n")) // write newline before enc
			qp.counter = encLen // reset counter after newline
		}
		// set counter to next line's char length 
		buf.Write(enc)
	}
	encoded.Write(buf.Next(1000))
	
} 

func (qp *QPEncoder) Encode(b []byte) ([]byte, error) {
	// split b by newlines
	qp.counter = 0
	var encoded bytes.Buffer
	for _, line := range bytes.Split(b, []byte("\n")) {
		qp.encodeLine(&encoded, &line)
	}
	
	return encoded.Bytes(), nil
}
