package adlib

type Record struct {
	PriRef                string          `xml:"priref,omitempty"`
	Institution           *Institution    `xml:"institution.name,omitempty"`
	Collection            string          `xml:"collection>term,omitempty"`
	ObjectNumber          string          `xml:"object_number,omitempty"`
	ObjectName            string          `xml:"Object_name>object_name>term,omitempty"`
	Title                 string          `xml:"Title>title,omitempty"`
	Description           string          `xml:"Description>description,omitempty"`
	Production            *Production     `xml:"Production,omitempty"`
	ProductionPeriod      []string        `xml:"production.period>term,omitempty"`
	ProductionDate        *ProductionDate `xml:"Production_date,omitempty"`
	Dimension             []Dimension     `xml:"Dimension,omitempty"`
	Material              []string        `xml:"Material>material>term,omitempty"`
	Taxonomy              *Taxonomy       `xml:"Taxonomy,omitempty"`
	FieldCollPlace        string          `xml:"field_coll.place>term,omitempty"`
	FieldCollPlaceFeature string          `xml:"field_coll.place.feature>term,omitempty"`
}

// TODO address.place,institution_code
type Institution struct {
	Name string `xml:"name,omitempty"`
}

// TODO creator
type Production struct {
	Place string `xml:"production.place>term,omitempty"`
}

type ProductionDate struct {
	Start string `xml:"production.date.start,omitempty"`
	End   string `xml:"production.date.end,omitempty"`
}

// TODO precision
type Dimension struct {
	Part  string `xml:"dimension.part,omitempty"`
	Type  string `xml:"dimension.type>term,omitempty"`
	Value string `xml:"dimension.value,omitempty"`
	Unit  string `xml:"dimension.unit>term,omitempty"`
}

type Taxonomy struct {
	CommonName string       `xml:"taxonomy.common_name,omitempty"`
	Rank       TaxonomyRank `xml:"taxonomy.rank,omitempty"`
}

type TaxonomyRank struct {
	Option string   `xml:"option,attr,omitempty"`
	Value  string   `xml:"value,attr,omitempty"`
	Text   []string `xml:"text,omitempty"`
}
