package deal

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
)
import "speedup/conf"


type Praser struct {
	Acction string
	Url string
	data []byte
}

func (p *Praser) Get_BytePackge() ([]byte, []byte)   {
	head := make([]byte,conf.HeadLen)
	body := p.Encode()
	bodyLen :=uint16(len(body))
	binary.BigEndian.PutUint16(head,bodyLen)
	return head,body
}

func (p *Praser) Prase_BytePackge(reader io.Reader)  {
	head := make([]byte,conf.HeadLen)
	reader.Read(head)
	bodyLen := binary.BigEndian.Uint16(head)
	body := make([]byte,bodyLen)
	reader.Read(body)
	*p = *p.Decode(body)
}

func (p *Praser) Read(out []byte)(int,error)  {
	outLeng := len(out)
	head,body := p.Get_BytePackge()
	if outLeng == len(head){
		copy(out,head)
	}else if outLeng == len(body){
		copy(out,body)
	}else {
		newDat:=append(head,body...)
		copy(out,newDat)
		panic("get header with body")
		return 0,nil
	}
	return 0,nil
}


func (p *Praser) Decode(raw []byte) *Praser  {
	reader := bytes.NewReader(raw)
	newp := new(Praser)
	gob.NewDecoder(reader).Decode(newp)
	return newp
}

func (p *Praser) Encode() []byte {
	a := &bytes.Buffer{}
	gob.NewEncoder(a).Encode(p)
	return a.Bytes()
}
