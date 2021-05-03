package adlib

import (
	"encoding/xml"
	"io"
)

type groupedXMLDecoder struct {
	xmlDecoder *xml.Decoder
}

func NewGroupedXMLDecoder(r io.Reader) Decoder {
	dec := xml.NewDecoder(r)
	dec.Strict = false
	return &groupedXMLDecoder{dec}
}

func (d *groupedXMLDecoder) Decode() (*Record, error) {
	for {
		dec := d.xmlDecoder
		tok, err := dec.Token()
		if tok == nil || err == io.EOF {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

		switch ty := tok.(type) {
		case xml.StartElement:
			if ty.Name.Local == "record" {
				rec := Record{}
				if err = dec.DecodeElement(&rec, &ty); err != nil {
					return nil, err
				}
				return &rec, nil
			}
		}
	}
}
