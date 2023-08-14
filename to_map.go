package wktcrs

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fikin/wkt-crs-go/wktcrsv1"
)

func propsFileNodeToArr(node antlr.Tree) (ret interface{}, err error) {
	switch n := node.(type) {
	case wktcrsv1.IPropsFileContext:
		ret := []interface{}{}
		for _, v := range n.AllPropRow() {
			o, err := nodeToMap(v)
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

// nolint:gocognit
func nodeToMap(node antlr.Tree) (interface{}, error) {
	if node == nil {
		return nil, nil
	}
	switch n := node.(type) {
	case wktcrsv1.IPropRowContext:
		return nodeToMap(n.EpsgDefLine())
	case wktcrsv1.IEpsgDefLineContext:
		return nodeToMap(n.Wkt())
	case wktcrsv1.IWktContext:
		return createMap(mapKeysBuilder{
			"compdcs": func() (interface{}, error) { return nodeToMap(n.Compdcs()) },
			"projcs":  func() (interface{}, error) { return nodeToMap(n.Projcs()) },
			"geogcs":  func() (interface{}, error) { return nodeToMap(n.Geogcs()) },
			"vertcs":  func() (interface{}, error) { return nodeToMap(n.Vertcs()) },
			"geoccs":  func() (interface{}, error) { return nodeToMap(n.Geoccs()) },
			"localcs": func() (interface{}, error) { return nodeToMap(n.Localcs()) },
		})
	case wktcrsv1.ICompdcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"projcs":    func() (interface{}, error) { return nodeToMap(n.Projcs()) },
			"geogcs":    func() (interface{}, error) { return nodeToMap(n.Geogcs()) },
			"vertcs":    func() (interface{}, error) { return nodeToMap(n.Vertcs()) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IProjcsContext:
		return createMap(mapKeysBuilder{
			"name":       func() (interface{}, error) { return nodeToStr(n.Name()) },
			"geogcs":     func() (interface{}, error) { return nodeToMap(n.Geogcs()) },
			"projection": func() (interface{}, error) { return nodeToMap(n.Projection()) },
			"parameter":  func() (interface{}, error) { return nodeToArr(paramCastToTree(n.AllParameter())) },
			"unit":       func() (interface{}, error) { return nodeToMap(n.Unit()) },
			"axis":       func() (interface{}, error) { return nodeToArr(axisCastToTree(n.AllAxis())) },
			"authority":  func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IGeoccsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"datum":     func() (interface{}, error) { return nodeToMap(n.Datum()) },
			"primem":    func() (interface{}, error) { return nodeToMap(n.Primem()) },
			"unit":      func() (interface{}, error) { return nodeToMap(n.Unit()) },
			"axis":      func() (interface{}, error) { return nodeToArr(axisCastToTree(n.AllAxis())) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IGeogcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"datum":     func() (interface{}, error) { return nodeToMap(n.Datum()) },
			"primem":    func() (interface{}, error) { return nodeToMap(n.Primem()) },
			"unit":      func() (interface{}, error) { return nodeToMap(n.Unit()) },
			"axis":      func() (interface{}, error) { return nodeToArr(axisCastToTree(n.AllAxis())) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IVertcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"datum":     func() (interface{}, error) { return nodeToMap(n.Vertdatum()) },
			"unit":      func() (interface{}, error) { return nodeToMap(n.Unit()) },
			"axis":      func() (interface{}, error) { return nodeToMap(n.Axis()) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.ILocalcsContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"datum":     func() (interface{}, error) { return nodeToMap(n.Localdatum()) },
			"unit":      func() (interface{}, error) { return nodeToMap(n.Unit()) },
			"axis":      func() (interface{}, error) { return nodeToArr(axisCastToTree(n.AllAxis())) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IDatumContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"spheroid":  func() (interface{}, error) { return nodeToMap(n.Spheroid()) },
			"towgs84":   func() (interface{}, error) { return nodeToMap(n.Towgs84()) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IVertdatumContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"type":      func() (interface{}, error) { return nodeToNumber(n.Type_()) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.ILocaldatumContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"type":      func() (interface{}, error) { return nodeToNumber(n.Type_()) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.ISpheroidContext:
		return createMap(mapKeysBuilder{
			"name":              func() (interface{}, error) { return nodeToStr(n.Name()) },
			"semiMajorAxis":     func() (interface{}, error) { return nodeToNumber(n.SemiMajorAxis()) },
			"inverseFlattening": func() (interface{}, error) { return nodeToNumber(n.InverseFlattening()) },
			"authority":         func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.ITowgs84Context:
		return createMap(mapKeysBuilder{
			"Dx": func() (interface{}, error) { return nodeToNumber(n.DXBF()) },
			"Dy": func() (interface{}, error) { return nodeToNumber(n.DYBF()) },
			"Dz": func() (interface{}, error) { return nodeToNumber(n.DZBF()) },
			"Rx": func() (interface{}, error) { return nodeToNumber(n.RXBF()) },
			"Ry": func() (interface{}, error) { return nodeToNumber(n.RYBF()) },
			"Rz": func() (interface{}, error) { return nodeToNumber(n.RZBF()) },
			"M":  func() (interface{}, error) { return nodeToNumber(n.MBF()) },
		})
	case wktcrsv1.IAuthorityContext:
		return createMap(mapKeysBuilder{
			"name": func() (interface{}, error) { return nodeToStr(n.AuthorityName()) },
			"code": func() (interface{}, error) { return nodeToStr(n.Code()) },
		})
	case wktcrsv1.IPrimemContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"longitude": func() (interface{}, error) { return nodeToNumber(n.Longitude()) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IUnitContext:
		return createMap(mapKeysBuilder{
			"name":        func() (interface{}, error) { return nodeToStr(n.Name()) },
			"angularUnit": func() (interface{}, error) { return nodeToNumber(n.AngularUnit()) },
		})
	case wktcrsv1.IAxisContext:
		return createMap(mapKeysBuilder{
			"name": func() (interface{}, error) { return nodeToStr(n.Name()) },
			// nolint:unparam
			"orientation": func() (interface{}, error) { return n.AxisOrient().GetText(), nil },
		})
	case wktcrsv1.IProjectionContext:
		return createMap(mapKeysBuilder{
			"name":      func() (interface{}, error) { return nodeToStr(n.Name()) },
			"authority": func() (interface{}, error) { return nodeToMap(n.Authority()) },
		})
	case wktcrsv1.IParameterContext:
		return createMap(mapKeysBuilder{
			"name":  func() (interface{}, error) { return nodeToStr(n.Name()) },
			"value": func() (interface{}, error) { return nodeToNumber(n.Value()) },
		})
	default:
		return nil, fmt.Errorf("expected object node type here but found : %s : %v", n, node)
	}
}

func nodeToStr(node antlr.Tree) (interface{}, error) {
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
	str = str[1 : len(str)-1]
	return str, nil
}

func nodeToNumber(node antlr.Tree) (interface{}, error) {
	if node == nil {
		return nil, nil
	}
	var str string
	switch n := node.(type) {
	case wktcrsv1.ITypeContext,
		wktcrsv1.ISemiMajorAxisContext,
		wktcrsv1.IInverseFlatteningContext,
		wktcrsv1.IDXBFContext,
		wktcrsv1.IDYBFContext,
		wktcrsv1.IDZBFContext,
		wktcrsv1.IRXBFContext,
		wktcrsv1.IRYBFContext,
		wktcrsv1.IRZBFContext,
		wktcrsv1.IMBFContext,
		wktcrsv1.ILongitudeContext,
		wktcrsv1.IAngularUnitContext,
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
	return f, nil
}

type mapObj map[string]interface{}

func nodeToArr(nodes []antlr.Tree) (interface{}, error) {
	arr := []mapObj{}
	for _, n := range nodes {
		o, err := nodeToMap(n)
		if err != nil {
			return nil, err
		}
		if o != nil {
			oo, ok := o.(mapObj)
			if !ok {
				return nil, fmt.Errorf("failed typcasting to jsonObj : %v", o)
			}
			arr = append(arr, oo)
		}
	}
	return arr, nil
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

type mapKeysBuilder map[string]func() (interface{}, error)

func createMap(props mapKeysBuilder) (mapObj, error) {
	ret := mapObj{}
	for k, v := range props {
		p, err := v()
		if err != nil {
			return nil, fmt.Errorf("%s : %#w", k, err)
		}
		if p != nil {
			ret[k] = p
		}
	}
	return ret, nil
}
