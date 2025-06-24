// Package gogpkgwkb provides utilitites to extract a WKB (Well-Known Binary) geometry from Geopackage Geometry data.
package gogpkgwkb

const (
	fixedHeaderSize uint8 = 8
	flagPosition    uint8 = 3
)

func getEnvelopeSize(flag byte) uint8 {
	envelopeMap := map[uint8]uint8{
		0: 0,
		1: 32,
		2: 48,
		3: 48,
		4: 64,
	}
	envelopeCode := (flag >> 1) & 0b111
	return fixedHeaderSize + envelopeMap[envelopeCode]
}

func ExtractWKB(rawGpkgGeom []byte) []byte {
	headerSize := getEnvelopeSize(rawGpkgGeom[flagPosition])
	return rawGpkgGeom[headerSize:]
}
