package adlib

type Decoder interface {
	Decode() (*Record, error)
}
