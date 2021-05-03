package adlib

type Record struct {
	PriRef                string          `xml:"priref,omitempty" json:"priRef,omitempty"`
	Institution           *Institution    `xml:"institution.name,omitempty" json:"institution,omitempty"`
	Collection            string          `xml:"collection>term,omitempty" json:"collection,omitempty"`
	ObjectCategory        string          `xml:"object_category>term,omitempty" json:"objectCategory,omitempty"`
	ObjectNumber          string          `xml:"object_number,omitempty" json:"objectNumber,omitempty"`
	ObjectName            string          `xml:"Object_name>object_name>term,omitempty" json:"objectName,omitempty"`
	Title                 string          `xml:"Title>title,omitempty" json:"title,omitempty"`
	Description           string          `xml:"Description>description,omitempty" json:"description,omitempty"`
	Production            *Production     `xml:"Production,omitempty" json:"production,omitempty"`
	ProductionPeriod      []string        `xml:"production.period>term,omitempty" json:"productionPeriod,omitempty"`
	ProductionDate        *ProductionDate `xml:"Production_date,omitempty" json:"productionDate,omitempty"`
	Dimension             []Dimension     `xml:"Dimension,omitempty" json:"dimension,omitempty"`
	Material              []string        `xml:"Material>material>term,omitempty" json:"material,omitempty"`
	Taxonomy              *Taxonomy       `xml:"Taxonomy,omitempty" json:"taxonomy,omitempty"`
	FieldCollPlace        []string        `xml:"field_coll.place>term,omitempty" json:"fieldCollPlace,omitempty"`
	FieldCollPlaceFeature []string        `xml:"field_coll.place.feature>term,omitempty" json:"fieldCollPlaceFeature,omitempty"`
}

// TODO address.place,institution_code
type Institution struct {
	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

// TODO creator.role
type Production struct {
	Creator []Creator `xml:"creator,omitempty" json:"creator,omitempty"`
	Place   string    `xml:"production.place>term,omitempty" json:"place,omitempty"`
}

// TODO biography
type Creator struct {
	Name      string `xml:"name,omitempty" json:"name,omitempty"`
	BirthDate string `xml:"birth.date.start,omitempty" json:"birthDate,omitempty"`
	DeathDate string `xml:"death.date.start,omitempty" json:"deathDate,omitempty"`
}

// TODO production.date.start.prec, production.date.end.prec
type ProductionDate struct {
	Start string `xml:"production.date.start,omitempty" json:"start,omitempty"`
	End   string `xml:"production.date.end,omitempty" json:"end,omitempty"`
}

// TODO precision
type Dimension struct {
	Part  string `xml:"dimension.part,omitempty" json:"part,omitempty"`
	Type  string `xml:"dimension.type>term,omitempty" json:"type,omitempty"`
	Value string `xml:"dimension.value,omitempty" json:"value,omitempty"`
	Unit  string `xml:"dimension.unit>term,omitempty" json:"unit,omitempty"`
}

type Taxonomy struct {
	CommonName string       `xml:"taxonomy.common_name,omitempty" json:"commonName,omitempty"`
	Rank       TaxonomyRank `xml:"taxonomy.rank,omitempty" json:"rank,omitempty"`
}

type TaxonomyRank struct {
	Option string   `xml:"option,attr,omitempty" json:"option,omitempty"`
	Value  string   `xml:"value,attr,omitempty" json:"value,omitempty"`
	Text   []string `xml:"text,omitempty" json:"text,omitempty"`
}
