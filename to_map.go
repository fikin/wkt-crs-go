package wktcrs

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
)

// PropsFileNodeToArr is returning array of map[string]interface{} for
// each EPSG definition in the input AST.
// Input must be content of a properties file.
func PropsFileNodeToArr(node antlr.Tree) (ret []MapObj, err error) {
	switch n := node.(type) {
	case wktcrsv1.IPropsFileContext:
		ret := []MapObj{}
		for _, v := range n.AllPropRow() {
			o, err := NodeToMap(v)
			if err != nil {
				return nil, err
			}
			if o != nil {
				ret = append(ret, o)
			}
		}
		return ret, nil
	default:
		return nil, fmt.Errorf("expected tree node type here but found : %s : %v", n, node)
	}
}

// NodeToMap is converting given node into json-like hashmap object
// nolint:gocognit
func NodeToMap(node antlr.Tree) (MapObj, error) {
	if node == nil {
		return nil, nil
	}
	switch n := node.(type) {
	case wktcrsv1.IPropRowContext:
		return NodeToMap(n.EpsgDefLine())
	case wktcrsv1.IEpsgDefLineContext:
		return NodeToMap(n.Wkt())
	case wktcrsv1.IWktContext:
		return createMap(mapKeysBuilder{
			"compdcs": func() (any, error) { return NodeToMap(n.Compdcs()) },
			"projcs":  func() (any, error) { return NodeToMap(n.Projcs()) },
			"geogcs":  func() (any, error) { return NodeToMap(n.Geogcs()) },
			"vertcs":  func() (any, error) { return NodeToMap(n.Vertcs()) },
			"geoccs":  func() (any, error) { return NodeToMap(n.Geoccs()) },
			"localcs": func() (any, error) { return NodeToMap(n.Localcs()) },
		})
	case wktcrsv1.ICompdcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"projcs":    func() (any, error) { return NodeToMap(n.Projcs()) },
			"geogcs":    func() (any, error) { return NodeToMap(n.Geogcs()) },
			"vertcs":    func() (any, error) { return NodeToMap(n.Vertcs()) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IProjcsContext:
		return createMap(mapKeysBuilder{
			"name":       func() (any, error) { return nodeToStr(n.Name()) },
			"geogcs":     func() (any, error) { return NodeToMap(n.Geogcs()) },
			"projection": func() (any, error) { return NodeToMap(n.Projection()) },
			"parameter":  func() (any, error) { return NodeToArr(paramCastToTree(n.AllParameter())) },
			"unit":       func() (any, error) { return NodeToMap(n.Unit()) },
			"axis":       func() (any, error) { return NodeToArr(axisCastToTree(n.AllAxis())) },
			"authority":  func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IGeoccsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"datum":     func() (any, error) { return NodeToMap(n.Datum()) },
			"primem":    func() (any, error) { return NodeToMap(n.Primem()) },
			"unit":      func() (any, error) { return NodeToMap(n.Unit()) },
			"axis":      func() (any, error) { return NodeToArr(axisCastToTree(n.AllAxis())) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IGeogcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"datum":     func() (any, error) { return NodeToMap(n.Datum()) },
			"primem":    func() (any, error) { return NodeToMap(n.Primem()) },
			"unit":      func() (any, error) { return NodeToMap(n.Unit()) },
			"axis":      func() (any, error) { return NodeToArr(axisCastToTree(n.AllAxis())) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IVertcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"datum":     func() (any, error) { return NodeToMap(n.Vertdatum()) },
			"unit":      func() (any, error) { return NodeToMap(n.Unit()) },
			"axis":      func() (any, error) { return NodeToMap(n.Axis()) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.ILocalcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"datum":     func() (any, error) { return NodeToMap(n.Localdatum()) },
			"unit":      func() (any, error) { return NodeToMap(n.Unit()) },
			"axis":      func() (any, error) { return NodeToArr(axisCastToTree(n.AllAxis())) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IDatumContext:
		return createMap(mapKeysBuilder{
			"name":     func() (any, error) { return nodeToStr(n.Name()) },
			"spheroid": func() (any, error) { return NodeToMap(n.Spheroid()) },
			"towgs84": func() (any, error) {
				return NodeToMap(n.Towgs84())
			},
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IVertdatumContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"type":      func() (any, error) { return nodeToNumber(n.Type_()) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.ILocaldatumContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"type":      func() (any, error) { return nodeToNumber(n.Type_()) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.ISpheroidContext:
		return createMap(mapKeysBuilder{
			"name":              func() (any, error) { return nodeToStr(n.Name()) },
			"semiMajorAxis":     func() (any, error) { return nodeToNumber(n.SemiMajorAxis()) },
			"inverseFlattening": func() (any, error) { return nodeToNumber(n.InverseFlattening()) },
			"authority":         func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.ITowgs84Context:
		return createMap(mapKeysBuilder{
			"Dx": func() (any, error) { return nodeToNumber(n.Dx()) },
			"Dy": func() (any, error) { return nodeToNumber(n.Dy()) },
			"Dz": func() (any, error) { return nodeToNumber(n.Dz()) },
			"Rx": func() (any, error) { return nodeToNumber(n.Ex()) },
			"Ry": func() (any, error) { return nodeToNumber(n.Ey()) },
			"Rz": func() (any, error) { return nodeToNumber(n.Ez()) },
			"M":  func() (any, error) { return nodeToNumber(n.Ppm()) },
		})
	case wktcrsv1.IAuthorityContext:
		return createMap(mapKeysBuilder{
			"name": func() (any, error) { return nodeToStr(n.AuthorityName()) },
			"code": func() (any, error) { return nodeToStr(n.Code()) },
		})
	case wktcrsv1.IPrimemContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"longitude": func() (any, error) { return nodeToNumber(n.Longitude()) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IUnitContext:
		return createMap(mapKeysBuilder{
			"name":        func() (any, error) { return nodeToStr(n.Name()) },
			"angularUnit": func() (any, error) { return nodeToNumber(n.ConversionFactor()) },
		})
	case wktcrsv1.IAxisContext:
		return createMap(mapKeysBuilder{
			"name": func() (any, error) { return nodeToStr(n.Name()) },
			// nolint:unparam
			"orientation": func() (any, error) { return n.AxisOrient().GetText(), nil },
		})
	case wktcrsv1.IProjectionContext:
		return createMap(mapKeysBuilder{
			"name":      func() (any, error) { return nodeToStr(n.Name()) },
			"authority": func() (any, error) { return NodeToMap(n.Authority()) },
		})
	case wktcrsv1.IParameterContext:
		return createMap(mapKeysBuilder{
			"name":  func() (any, error) { return nodeToStr(n.Name()) },
			"value": func() (any, error) { return nodeToNumber(n.Value()) },
		})
	default:
		return nil, fmt.Errorf("expected object node type here but found : %s : %#v", n, node)
	}
}

func nodeToStr(node antlr.Tree) (*string, error) {
	if node == nil {
		return nil, nil
	}
	var str string
	switch n := node.(type) {
	case wktcrsv1.IAxisOrientContext,
		wktcrsv1.IAuthorityNameContext,
		wktcrsv1.INameContext,
		wktcrsv1.ICodeContext:
		str = n.(antlr.ParseTree).GetText()
	default:
		return nil, fmt.Errorf("expected string node type here but found : %s : %v", n, node)
	}
	str = stripQuotes(str)
	return &str, nil
}

func stripQuotes(str string) string {
	if str[0:1] == "\"" {
		return str[1 : len(str)-1]
	}
	return str
}

func nodeToNumber(node antlr.Tree) (*float64, error) {
	if node == nil {
		return nil, nil
	}
	var str string
	switch n := node.(type) {
	case wktcrsv1.ITypeContext,
		wktcrsv1.ISemiMajorAxisContext,
		wktcrsv1.IInverseFlatteningContext,
		wktcrsv1.IDxContext,
		wktcrsv1.IDyContext,
		wktcrsv1.IDzContext,
		wktcrsv1.IExContext,
		wktcrsv1.IEyContext,
		wktcrsv1.IEzContext,
		wktcrsv1.IPpmContext,
		wktcrsv1.ILongitudeContext,
		wktcrsv1.IConversionFactorContext,
		wktcrsv1.IValueContext,
		wktcrsv1.INumberContext:
		str = n.(antlr.ParseTree).GetText()
	default:
		return nil, fmt.Errorf("expected number node type here but found : %s : %v", n, node)
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

// MapObj is shortcut to type a json-like object as hashmap
type MapObj map[string]interface{}

// NodeToArr is converting given nodes into array of objects.
func NodeToArr(nodes []antlr.Tree) ([]MapObj, error) {
	arr := []MapObj{}
	for _, n := range nodes {
		o, err := NodeToMap(n)
		if err != nil {
			return nil, err
		}
		if o != nil {
			arr = append(arr, o)
		}
	}
	if len(arr) > 0 {
		return arr, nil
	}
	return nil, nil
}

func axisCastToTree(nodes []wktcrsv1.IAxisContext) []antlr.Tree {
	arr := make([]antlr.Tree, len(nodes))
	for i, v := range nodes {
		arr[i] = v
	}
	return arr
}

func paramCastToTree(nodes []wktcrsv1.IParameterContext) []antlr.Tree {
	arr := make([]antlr.Tree, len(nodes))
	for i, v := range nodes {
		arr[i] = v
	}
	return arr
}

type mapKeysBuilder map[string]func() (any, error)

func createMap(props mapKeysBuilder) (MapObj, error) {
	notEmpty := false
	ret := MapObj{}
	for k, v := range props {
		p, err := v()
		if err != nil {
			return nil, fmt.Errorf("%s : %#w", k, err)
		}
		if !isNilFixed(p) {
			ret[k] = p
			notEmpty = true
		}
	}
	if notEmpty {
		return ret, nil
	}
	return nil, nil
}

func isNilFixed(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
