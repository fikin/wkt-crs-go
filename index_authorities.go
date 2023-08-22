package wktcrs

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type DatumDef struct {
	Name      string             `json:"name"`
	Type      string             `json:"type"`
	Ellipsoid *string            `json:"ellipsoid"`
	ToWGS84   map[string]float64 `json:"toWgs84,omitempty"`
	Node      antlr.Tree         `json:"-"`
}

type EllipsoidDef struct {
	Name              string     `json:"name"`
	SemiMajorAxis     float64    `json:"semiMajorAxis"`
	InverseFlattening float64    `json:"inverseFlattening"`
	Node              antlr.Tree `json:"-"`
}

type GeoGcsDef struct {
	Name  string     `json:"name"`
	Datum string     `json:"datum"`
	Node  antlr.Tree `json:"-"`
}

type GeoCcsDef struct {
	Name  string     `json:"name"`
	Datum string     `json:"datum"`
	Node  antlr.Tree `json:"-"`
}

type ProjcsDef struct {
	Name       string             `json:"name"`
	GeoGcs     string             `json:"geogcs"`
	Projection string             `json:"projection"`
	Params     map[string]float64 `json:"params,omitempty"`
	Node       antlr.Tree         `json:"-"`
}

type Vertcs struct {
	Name  string     `json:"name"`
	Datum string     `json:"datum"`
	Node  antlr.Tree `json:"-"`
}

type Localcs struct {
	Name  string     `json:"name"`
	Datum string     `json:"datum"`
	Node  antlr.Tree `json:"-"`
}

type Compdcs struct {
	Name     string     `json:"name"`
	FirstCs  string     `json:"firstcs,omitempty"`
	SecondCs string     `json:"secondcs,omitempty"`
	Node     antlr.Tree `json:"-"`
}

type AuthoritiesIndx struct {
	Authorities map[string]string       `json:"authorities,omitempty"`
	Projections map[string]string       `json:"projections,omitempty"`
	Datums      map[string]DatumDef     `json:"datums,omitempty"`
	Ellipshoids map[string]EllipsoidDef `json:"ellipsoids,omitempty"`
	GeoGcs      map[string]GeoGcsDef    `json:"geographiccs,omitempty"`
	GeoCcs      map[string]GeoCcsDef    `json:"geocentriccs,omitempty"`
	Projcs      map[string]ProjcsDef    `json:"projectioncs,omitempty"`
	Vertcs      map[string]Vertcs       `json:"verticalcs,omitempty"`
	Localcs     map[string]Localcs      `json:"localcs,omitempty"`
	Compdcs     map[string]Compdcs      `json:"compdcs,omitempty"`
}

// AuthoritiesIndexPropFile is indexing a repo of definitions
func AuthoritiesIndexPropFile(node wktcrsv1.IPropsFileContext) (*AuthoritiesIndx, error) {
	indx := &AuthoritiesIndx{
		Authorities: map[string]string{},
		Projections: map[string]string{},
		Datums:      map[string]DatumDef{},
		Ellipshoids: map[string]EllipsoidDef{},
		GeoGcs:      map[string]GeoGcsDef{},
		GeoCcs:      map[string]GeoCcsDef{},
		Projcs:      map[string]ProjcsDef{},
		Vertcs:      map[string]Vertcs{},
		Localcs:     map[string]Localcs{},
		Compdcs:     map[string]Compdcs{},
	}
	err := authoritiesIndexNode(indx, node)
	return indx, err
}

func authoritiesIndexNode(indx *AuthoritiesIndx, node antlr.Tree) error {
	switch n := node.(type) {
	case wktcrsv1.ICompdcsContext:
		_, err := authorithesIndxCompdcs(indx, n)
		return err
	case wktcrsv1.IVertcsContext:
		_ = authoritiesIndexVertcs(indx, n)
		return nil
	case wktcrsv1.ILocalcsContext:
		_ = authoritiesIndexLocalcs(indx, n)
		return nil
	case wktcrsv1.IProjcsContext:
		_, err := authoritiesIndexProjcs(indx, n)
		return err
	case wktcrsv1.IGeoccsContext:
		_ = authoritiesIndexGeoccs(indx, n)
		return nil
	case wktcrsv1.IGeogcsContext:
		_ = authoritiesIndexGeogcs(indx, n)
		return nil
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
		return visitChildren(n, func(t antlr.Tree) error {
			return authoritiesIndexNode(indx, t)
		})
	}
}

func authorithesIndxCompdcs(indx *AuthoritiesIndx, n wktcrsv1.ICompdcsContext) (string, error) {
	var err error
	var first string
	switch {
	case n.Projcs() != nil:
		first, err = authoritiesIndexProjcs(indx, n.Projcs())
		if err != nil {
			return "", err
		}
	case n.Geogcs() != nil:
		first = authoritiesIndexGeogcs(indx, n.Geogcs())
	default:
		return "", fmt.Errorf("compdcs has no first coordinate system : %v", n)
	}
	def := Compdcs{
		Name:     stripQuotes(n.Name().GetText()),
		FirstCs:  first,
		SecondCs: authoritiesIndexVertcs(indx, n.Vertcs()),
		Node:     n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	if old, ok := indx.Compdcs[authCode]; ok {
		if old.FirstCs == def.FirstCs && old.SecondCs == def.SecondCs {
			return authCode, nil
		}
		authCode = findNextFreeKey(maps.Keys(indx.GeoCcs), authCode)
	}
	indx.Compdcs[authCode] = def
	return authCode, nil
}

func authoritiesIndexLocalcs(indx *AuthoritiesIndx, n wktcrsv1.ILocalcsContext) string {
	def := Localcs{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: authoritiesIndexOtherDatum(indx, n.Localdatum().Authority(), n.Localdatum().Name(), n.Localdatum().Type_(), n.Localdatum()),
		Node:  n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	if old, ok := indx.Localcs[authCode]; ok {
		if old.Datum == def.Datum {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(indx.GeoCcs), authCode)
	}
	indx.Localcs[authCode] = def
	return authCode
}

func authoritiesIndexVertcs(indx *AuthoritiesIndx, n wktcrsv1.IVertcsContext) string {
	def := Vertcs{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: authoritiesIndexOtherDatum(indx, n.Vertdatum().Authority(), n.Vertdatum().Name(), n.Vertdatum().Type_(), n.Vertdatum()),
		Node:  n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	if old, ok := indx.Vertcs[authCode]; ok {
		if old.Datum == def.Datum {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(indx.GeoCcs), authCode)
	}
	indx.Vertcs[authCode] = def
	return authCode
}

func authoritiesIndexOtherDatum(indx *AuthoritiesIndx, auth wktcrsv1.IAuthorityContext, name wktcrsv1.INameContext, datumType wktcrsv1.ITypeContext, n antlr.Tree) string {
	def := DatumDef{
		Name: stripQuotes(name.GetText()),
		Type: stripQuotes(datumType.GetText()),
		Node: n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, auth, def.Name)
	return authoritiesIndexDatumID(indx, def, authCode)
}

func authoritiesIndexProjcs(indx *AuthoritiesIndx, n wktcrsv1.IProjcsContext) (string, error) {
	params, err := projParamsToMap(n)
	if err != nil {
		return "", err
	}
	def := ProjcsDef{
		Name:       stripQuotes(n.Name().GetText()),
		GeoGcs:     authoritiesIndexGeogcs(indx, n.Geogcs()),
		Projection: authoritiesIndexProjection(indx, n.Projection()),
		Params:     params,
		Node:       n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	repo := indx.Projcs
	if old, ok := repo[authCode]; ok {
		if old.GeoGcs == def.GeoGcs && old.Projection == def.Projection && maps.Equal(old.Params, def.Params) {
			return authCode, nil
		}
		authCode = findNextFreeKey(maps.Keys(repo), authCode)
		indx.Authorities[authCode] = "UNDEF"
	}
	repo[authCode] = def
	return authCode, nil
}

func authoritiesIndexProjection(indx *AuthoritiesIndx, n wktcrsv1.IProjectionContext) string {
	name := textToCapitalized(stripQuotes(n.Name().GetText()))
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), name)
	repo := indx.Projections
	for k, v := range repo {
		if v == authCode {
			return k
		}
	}
	if old, ok := repo[authCode]; ok {
		if old == name {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(repo), authCode)
		indx.Authorities[authCode] = "UNDEF"
	}
	repo[authCode] = name
	return authCode
}

func authoritiesIndexGeoccs(indx *AuthoritiesIndx, n wktcrsv1.IGeoccsContext) string {
	def := GeoCcsDef{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: authoritiesIndexDatum(indx, n.Datum()),
		Node:  n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	repo := indx.GeoCcs
	if old, ok := repo[authCode]; ok {
		if old.Datum == def.Datum {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(repo), authCode)
		indx.Authorities[authCode] = "UNDEF"
	}
	repo[authCode] = def
	return authCode
}

func authoritiesIndexGeogcs(indx *AuthoritiesIndx, n wktcrsv1.IGeogcsContext) string {
	def := GeoGcsDef{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: authoritiesIndexDatum(indx, n.Datum()),
		Node:  n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	repo := indx.GeoGcs
	if old, ok := repo[authCode]; ok {
		if old.Datum == def.Datum {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(repo), authCode)
		indx.Authorities[authCode] = "UNDEF"
	}
	repo[authCode] = def
	return authCode
}

func authoritiesIndexDatum(indx *AuthoritiesIndx, n wktcrsv1.IDatumContext) string {
	towgs, err := wgs84ToMap(n.Towgs84())
	if err != nil {
		return ""
	}
	ellipsoid, err := authoritiesIndexSpheroidCode(indx, n.Spheroid())
	if err != nil {
		return ""
	}
	def := DatumDef{
		Name:      stripQuotes(n.Name().GetText()),
		Type:      "datum",
		Ellipsoid: ellipsoid,
		ToWGS84:   towgs,
		Node:      n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	return authoritiesIndexDatumID(indx, def, authCode)
}

func authoritiesIndexDatumID(indx *AuthoritiesIndx, def DatumDef, authCode string) string {
	if old, ok := indx.Datums[authCode]; ok {
		if old.Type == def.Type &&
			strPtrEqual(old.Ellipsoid, def.Ellipsoid) &&
			maps.Equal(old.ToWGS84, def.ToWGS84) {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(indx.Datums), authCode)
	}
	indx.Datums[authCode] = def
	return authCode
}

func strPtrEqual(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		return *a == *b
	}
	return false
}

func authoritiesIndexSpheroidCode(indx *AuthoritiesIndx, n wktcrsv1.ISpheroidContext) (*string, error) {
	if n == nil {
		return nil, nil
	}
	arr, err := nodesToNumbers([]string{"A", "B"}, []antlr.Tree{n.SemiMajorAxis(), n.InverseFlattening()})
	if err != nil {
		return nil, err
	}
	def := EllipsoidDef{
		Name:              textToCapitalized(stripQuotes(n.Name().GetText())),
		SemiMajorAxis:     arr["A"],
		InverseFlattening: arr["B"],
		Node:              n,
	}
	authCode := authoritiesIndexAuthorityCode(indx, n.Authority(), def.Name)
	repo := indx.Ellipshoids
	if old, ok := repo[authCode]; ok {
		if old.InverseFlattening == def.InverseFlattening && old.SemiMajorAxis == def.SemiMajorAxis {
			return &authCode, nil
		}
		authCode = findNextFreeKey(maps.Keys(repo), authCode)
		indx.Authorities[authCode] = "UNDEF"
	}
	repo[authCode] = def
	return &authCode, nil
}

func findNextFreeKey(tbl []string, key string) string {
	slices.Sort(tbl)
	i := 2
	for {
		k := fmt.Sprintf("%s_%d", key, i)
		if !slices.Contains(tbl, k) {
			return k
		}
		i++
	}
}

func authoritiesIndexAuthorityCode(indx *AuthoritiesIndx, n wktcrsv1.IAuthorityContext, name string) string {
	if n == nil {
		authCode := textToCapitalized(name)
		if _, ok := indx.Authorities[authCode]; ok {
			return authCode
		}
		indx.Authorities[authCode] = "UNDEF"
		return name
	}
	authCode := stripQuotes(n.Code().GetText())
	authName := stripQuotes(n.AuthorityName().GetText())
	if old, ok := indx.Authorities[authCode]; ok {
		if old != authName {
			authCode = findNextFreeKey(maps.Keys(indx.Authorities), authCode)
		}
	}
	indx.Authorities[authCode] = authName
	return authCode
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9_ ]+`)

func textToCapitalized(str string) string {
	return textToSomething(str, " ")
}

func textToCammel(str string) string {
	return textToSomething(str, "")
}

func textToSomething(str, join string) string {
	arr := strings.Split(strings.ReplaceAll(nonAlphanumericRegex.ReplaceAllString(strings.ToLower(str), ""), "_", " "), " ")
	for i, s := range arr {
		// nolint:staticcheck
		arr[i] = strings.Title(s)
	}
	return strings.Join(arr, join)
}

func projParamsToMap(node antlr.Tree) (map[string]float64, error) {
	ret := map[string]float64{}
	for _, nn := range node.GetChildren() {
		// nolint:gocritic
		switch n := nn.(type) {
		case wktcrsv1.IParameterContext:
			val, err := nodeToNumber(n.Value())
			if err != nil {
				return nil, err
			}
			ret[textToCammel(stripQuotes(n.Name().GetText()))] = *val
		}
	}
	return ret, nil
}

func wgs84ToMap(node wktcrsv1.ITowgs84Context) (map[string]float64, error) {
	if node == nil {
		return nil, nil
	}
	arr, err := nodesToNumbers(
		[]string{"Dx", "Dy", "Dz", "Rx", "Ry", "Rz", "Mpp"},
		[]antlr.Tree{
			node.Dx(), node.Dy(), node.Dz(),
			node.Ex(), node.Ey(), node.Ez(),
			node.Ppm(),
		})
	return arr, err
}

func nodesToNumbers(keys []string, values []antlr.Tree) (map[string]float64, error) {
	arr := map[string]float64{}
	return arr, nodesToMapNumbers(arr, keys, values)
}

func nodesToMapNumbers(arr map[string]float64, keys []string, values []antlr.Tree) error {
	for i, n := range values {
		val, err := nodeToNumber(n)
		if err != nil {
			return err
		}
		if val != nil {
			arr[keys[i]] = *val
		}
	}
	return nil
}
