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
	Name      string              `json:"name"`
	Type      string              `json:"type"`
	Ellipsoid *string             `json:"ellipsoid"`
	ToWGS84   map[string]*float64 `json:"toWgs84,omitempty"`
	Node      antlr.Tree          `json:"-"`
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
		_, err := getCompdcs(indx, n)
		return err
	case wktcrsv1.IVertcsContext:
		_ = getVertcs(indx, n)
		return nil
	case wktcrsv1.ILocalcsContext:
		_ = getLocalcs(indx, n)
		return nil
	case wktcrsv1.IProjcsContext:
		_, err := getProjcs(indx, n)
		return err
	case wktcrsv1.IGeoccsContext:
		_ = getGeoccs(indx, n)
		return nil
	case wktcrsv1.IGeogcsContext:
		_ = getGeogcs(indx, n)
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

func getCompdcs(indx *AuthoritiesIndx, n wktcrsv1.ICompdcsContext) (string, error) {
	var err error
	var first string
	switch {
	case n.Projcs() != nil:
		first, err = getProjcs(indx, n.Projcs())
		if err != nil {
			return "", err
		}
	case n.Geogcs() != nil:
		first = getGeogcs(indx, n.Geogcs())
	default:
		return "", fmt.Errorf("compdcs has no first coordinate system : %v", n)
	}
	def := Compdcs{
		Name:     stripQuotes(n.Name().GetText()),
		FirstCs:  first,
		SecondCs: getVertcs(indx, n.Vertcs()),
		Node:     n,
	}
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
	if old, ok := indx.Compdcs[authCode]; ok {
		if old.FirstCs == def.FirstCs && old.SecondCs == def.SecondCs {
			return authCode, nil
		}
		authCode = findNextFreeKey(maps.Keys(indx.GeoCcs), authCode)
	}
	indx.Compdcs[authCode] = def
	return authCode, nil
}

func getLocalcs(indx *AuthoritiesIndx, n wktcrsv1.ILocalcsContext) string {
	def := Localcs{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: getOtherDatum(indx, n.Localdatum().Authority(), n.Localdatum().Name(), n.Localdatum().Type_(), n.Localdatum()),
		Node:  n,
	}
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
	if old, ok := indx.Localcs[authCode]; ok {
		if old.Datum == def.Datum {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(indx.GeoCcs), authCode)
	}
	indx.Localcs[authCode] = def
	return authCode
}

func getVertcs(indx *AuthoritiesIndx, n wktcrsv1.IVertcsContext) string {
	def := Vertcs{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: getOtherDatum(indx, n.Vertdatum().Authority(), n.Vertdatum().Name(), n.Vertdatum().Type_(), n.Vertdatum()),
		Node:  n,
	}
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
	if old, ok := indx.Vertcs[authCode]; ok {
		if old.Datum == def.Datum {
			return authCode
		}
		authCode = findNextFreeKey(maps.Keys(indx.GeoCcs), authCode)
	}
	indx.Vertcs[authCode] = def
	return authCode
}

func getOtherDatum(indx *AuthoritiesIndx, auth wktcrsv1.IAuthorityContext, name wktcrsv1.INameContext, datumType wktcrsv1.ITypeContext, n antlr.Tree) string {
	def := DatumDef{
		Name: stripQuotes(name.GetText()),
		Type: stripQuotes(datumType.GetText()),
		Node: n,
	}
	authCode := getAuthorityCode(indx, auth, def.Name)
	return getDatumID(indx, def, authCode)
}

func getProjcs(indx *AuthoritiesIndx, n wktcrsv1.IProjcsContext) (string, error) {
	def := ProjcsDef{
		Name:       stripQuotes(n.Name().GetText()),
		GeoGcs:     getGeogcs(indx, n.Geogcs()),
		Projection: getProjection(indx, n.Projection()),
		Params:     make(map[string]float64),
		Node:       n,
	}
	err := paramsToProjectionDef(&def, n)
	if err != nil {
		return "", err
	}
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
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

func getProjection(indx *AuthoritiesIndx, n wktcrsv1.IProjectionContext) string {
	name := textToCapitalized(stripQuotes(n.Name().GetText()))
	authCode := getAuthorityCode(indx, n.Authority(), name)
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

func getGeoccs(indx *AuthoritiesIndx, n wktcrsv1.IGeoccsContext) string {
	def := GeoCcsDef{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: getDatum(indx, n.Datum()),
		Node:  n,
	}
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
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

func getGeogcs(indx *AuthoritiesIndx, n wktcrsv1.IGeogcsContext) string {
	def := GeoGcsDef{
		Name:  stripQuotes(n.Name().GetText()),
		Datum: getDatum(indx, n.Datum()),
		Node:  n,
	}
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
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

func getDatum(indx *AuthoritiesIndx, n wktcrsv1.IDatumContext) string {
	towgs, err := toWGS84Def(n.Towgs84())
	if err != nil {
		return ""
	}
	ellipsoid, err := getSpheroidCode(indx, n.Spheroid())
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
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
	return getDatumID(indx, def, authCode)
}

func getDatumID(indx *AuthoritiesIndx, def DatumDef, authCode string) string {
	if old, ok := indx.Datums[authCode]; ok {
		if old.Type == def.Type &&
			strPtrEqual(old.Ellipsoid, def.Ellipsoid) &&
			maps.EqualFunc(old.ToWGS84, def.ToWGS84, func(a, b *float64) bool {
				if a == nil && b == nil {
					return true
				}
				if a != nil && b != nil {
					return *a == *b
				}
				return false
			}) {
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

func getSpheroidCode(indx *AuthoritiesIndx, n wktcrsv1.ISpheroidContext) (*string, error) {
	if n == nil {
		return nil, nil
	}
	arr, err := nodesToNumbers([]antlr.Tree{n.SemiMajorAxis(), n.InverseFlattening()})
	if err != nil {
		return nil, err
	}
	def := EllipsoidDef{
		Name:              textToCapitalized(stripQuotes(n.Name().GetText())),
		SemiMajorAxis:     *arr[0],
		InverseFlattening: *arr[1],
		Node:              n,
	}
	authCode := getAuthorityCode(indx, n.Authority(), def.Name)
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

func getAuthorityCode(indx *AuthoritiesIndx, n wktcrsv1.IAuthorityContext, name string) string {
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

func paramsToProjectionDef(p *ProjcsDef, node antlr.Tree) error {
	for _, nn := range node.GetChildren() {
		// nolint:gocritic
		switch n := nn.(type) {
		case wktcrsv1.IParameterContext:
			val, err := nodeToNumber(n.Value())
			if err != nil {
				return err
			}
			p.Params[textToCammel(stripQuotes(n.Name().GetText()))] = *val
		}
	}
	return nil
}

func toWGS84Def(node wktcrsv1.ITowgs84Context) (map[string]*float64, error) {
	if node == nil {
		return nil, nil
	}
	arr, err := nodesToNumbers([]antlr.Tree{
		node.Dx(), node.Dy(), node.Dz(),
		node.Ex(), node.Ey(), node.Ez(),
		node.Ppm(),
	})
	return map[string]*float64{
		"Dx":  arr[0],
		"Dy":  arr[1],
		"Dz":  arr[2],
		"Rx":  arr[3],
		"Ry":  arr[4],
		"Rz":  arr[5],
		"Mpp": arr[6],
	}, err
}

func nodesToNumbers(lst []antlr.Tree) ([]*float64, error) {
	arr := make([]*float64, len(lst))
	for i, n := range lst {
		val, err := nodeToNumber(n)
		if err != nil {
			return arr, err
		}
		arr[i] = val
	}
	return arr, nil
}
