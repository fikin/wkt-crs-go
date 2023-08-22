package wktcrs

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
)

type Wgs84GeoCs struct {
	AuthCode string             `json:"authCode"`
	AuthName string             `json:"authName"`
	CsName   string             `json:"csName"`
	Datum    map[string]float64 `json:"datum"`
	Node     antlr.Tree         `json:"-"`
}

type Wgs84ProjCs struct {
	AuthCode       string             `json:"authCode"`
	AuthName       string             `json:"authName"`
	CsName         string             `json:"csName"`
	Datum          map[string]float64 `json:"datum"`
	Projection     map[string]float64 `json:"projection"`
	ProjectionName string             `json:"projectionName"`
	Node           antlr.Tree         `json:"-"`
}

type Wgs84Repo struct {
	GeoGcs      []Wgs84GeoCs            `json:"geogcs"`
	GeoCcs      []Wgs84GeoCs            `json:"geoccs"`
	ProjCs      []Wgs84ProjCs           `json:"projcs"`
	Unsupported map[string][]antlr.Tree `json:"-"`
}

// Wgs84PropFile is indexing a repo of definitions
func Wgs84PropFile(node wktcrsv1.IPropsFileContext) (*Wgs84Repo, error) {
	indx := &Wgs84Repo{
		GeoGcs:      []Wgs84GeoCs{},
		GeoCcs:      []Wgs84GeoCs{},
		ProjCs:      []Wgs84ProjCs{},
		Unsupported: map[string][]antlr.Tree{},
	}
	err := wgs84Node(indx, node)
	return indx, err
}

func wgs84Node(indx *Wgs84Repo, node antlr.Tree) error {
	switch n := node.(type) {
	case wktcrsv1.ICompdcsContext,
		wktcrsv1.IVertcsContext,
		wktcrsv1.ILocalcsContext:
		wgs84AddUnsupported(indx.Unsupported, "Unsupported coordinate system", n)
		return nil
	case wktcrsv1.IProjcsContext:
		return wgs84ProjCs(indx, n)
	case wktcrsv1.IGeoccsContext:
		return wgs84GeoCcs(indx, n)
	case wktcrsv1.IGeogcsContext:
		return wgs84GeoGcs(indx, n)
	case wktcrsv1.IDatumContext,
		wktcrsv1.IVertdatumContext,
		wktcrsv1.ILocaldatumContext,
		wktcrsv1.ISpheroidContext,
		wktcrsv1.IAuthorityContext,
		wktcrsv1.IProjectionContext,
		wktcrsv1.ITowgs84Context,
		wktcrsv1.IPrimemContext,
		wktcrsv1.IUnitContext,
		wktcrsv1.IAxisContext:
		return fmt.Errorf("should not be here : %v", n)
	default:
		return VisitChildren(n, func(t antlr.Tree) error {
			return wgs84Node(indx, t)
		})
	}
}

func wgs84GeoGcs(indx *Wgs84Repo, n wktcrsv1.IGeogcsContext) error {
	if n.Datum().Towgs84() == nil {
		wgs84AddUnsupported(indx.Unsupported, "Missing TOWGS84", n)
		return nil
	}
	datum, err := wgs84Datum(n.Datum().Towgs84(), n.Datum().Spheroid())
	if err != nil {
		return err
	}
	def := Wgs84GeoCs{
		AuthCode: stripQuotes(n.Authority().Code().GetText()),
		AuthName: stripQuotes(n.Authority().AuthorityName().GetText()),
		CsName:   "GEOGCS",
		Datum:    datum,
		Node:     n,
	}
	indx.GeoGcs = append(indx.GeoGcs, def)
	return nil
}

func wgs84GeoCcs(indx *Wgs84Repo, n wktcrsv1.IGeoccsContext) error {
	if n.Datum().Towgs84() == nil {
		wgs84AddUnsupported(indx.Unsupported, "Missing TOWGS84", n)
		return nil
	}
	datum, err := wgs84Datum(n.Datum().Towgs84(), n.Datum().Spheroid())
	if err != nil {
		return err
	}
	def := Wgs84GeoCs{
		AuthCode: stripQuotes(n.Authority().Code().GetText()),
		AuthName: stripQuotes(n.Authority().AuthorityName().GetText()),
		CsName:   "GEOCCS",
		Datum:    datum,
		Node:     n,
	}
	indx.GeoCcs = append(indx.GeoCcs, def)
	return nil
}

func wgs84ProjCs(indx *Wgs84Repo, n wktcrsv1.IProjcsContext) error {
	if n.Geogcs().Datum().Towgs84() == nil {
		wgs84AddUnsupported(indx.Unsupported, "Missing TOWGS84", n)
		return nil
	}
	projName := textToCapitalized(stripQuotes(n.Projection().Name().GetText()))
	switch projName {
	case "Popular Visualisation Pseudo Mercator",
		"Lambert Azimuthal Equal Area",
		"Lambert Conformal Conic 2sp",
		"Transverse Mercator",
		"Albers Conic Equal Area":
	default:
		wgs84AddUnsupported(indx.Unsupported, "Missing projection support : "+projName, n)
		return nil
	}
	datum, err := wgs84Datum(n.Geogcs().Datum().Towgs84(), n.Geogcs().Datum().Spheroid())
	if err != nil {
		return err
	}
	params, err := projParamsToMap(n)
	if err != nil {
		return err
	}
	def := Wgs84ProjCs{
		AuthCode:       stripQuotes(n.Authority().Code().GetText()),
		AuthName:       stripQuotes(n.Authority().AuthorityName().GetText()),
		CsName:         "PROJCS",
		Datum:          datum,
		Projection:     params,
		ProjectionName: projName,
		Node:           n,
	}
	indx.ProjCs = append(indx.ProjCs, def)
	return nil
}

func wgs84AddUnsupported(tbl map[string][]antlr.Tree, key string, item antlr.Tree) {
	arr, ok := tbl[key]
	if !ok {
		arr = []antlr.Tree{}
	}
	arr = append(arr, item)
	tbl[key] = arr
}

func wgs84Datum(towgs84 wktcrsv1.ITowgs84Context, spheroid wktcrsv1.ISpheroidContext) (map[string]float64, error) {
	datum, err := wgs84ToMap(towgs84)
	if err != nil {
		return nil, err
	}
	err = nodesToMapNumbers(
		datum,
		[]string{"SemiMajorAxis", "InverseFlattening"},
		[]antlr.Tree{spheroid.SemiMajorAxis(), spheroid.InverseFlattening()},
	)
	return datum, err
}
