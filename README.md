# wkt-crs-go

Parser for Well Known Text (WKT) Coordinate Reference System (CRS) in golang.

It supports version 1 (see references).

References:

- [WKT](https://en.wikipedia.org/wiki/Well-known_text_representation_of_coordinate_reference_systems)
- [WKT v1](https://github.com/geotools/geotools/blob/main/modules/library/opengis/src/main/java/org/opengis/referencing/doc-files/WKT.html)

Grammar itself is maintained at [ANTLR grammars](https://github.com/antlr/grammars-v4) repo;

The parser is generated by ANTLR for `Go` target.
There are no ANTLR visitor nor listener generated, as they are rather useless from code point of view.

There are two main entry rules:

- `propsFile` is accepting properties file like the one from [GeoTools epsg.properties](https://raw.githubusercontent.com/geotools/geotools/main/modules/plugin/epsg-wkt/src/main/resources/org/geotools/referencing/epsg/wkt/epsg.properties)
- `wkt` is accepting single definition WKT text

The parser can be used to read WKT and provide AST.

There are few AST visitors which can be used as basis for whatever else is required.

- `to_text.go` is simple identity print of AST and it can be used to strip whitespace
- `to_pretty_text.go` is converting AST to text indented multi-line format
- `to_map` is converting AST to array of objects which can be written as `json`
- `index_authorities.go` is converting AST to array of objects which can be written as `json`
- `wgs84_repo.go` is indexing only CRS definitions with TOWGS85 defined

## Usage as library

In the core of AST is ANTLR generated lexer and parser code.

Add this lib as dependency:

```shell
go get github.com/fikin/wkt-crs-go
```

Sketch some code:

```go
// in this example we'll parse epsg.properties file
is, err := antlr.NewFileStream("path to epsg.properties")
require.NoError(t, err)

// creates parser and lexer and returns root AST node
rootNode := wktcrs.NewWktcrsv1Parser(is)

// starting rule for the parser in this case is PropsFile.
// this typically would depend on the input from above.
visitAST(rootNode.PropsFile())
```

## Command line usage

This repo offers a command line tool to process WKT via the command line.

Add this lib as dependency:

```shell
go install github.com/fikin/wkt-crs-go/cmd/parser@latest
```

Currently there are 3 input options : WKT, properties or properties+templates.

### "wkt" option

WKT is accepting text and it would format it.

```shell
// will strip out all whitespace
echo 'PROJCS[...' | ~/go/bin/parser wkt

// will pretty print
echo 'PROJCS[...' | ~/go/bin/parser wkt -format pretty

// will print it as json
echo 'PROJCS[...' | ~/go/bin/parser wkt -format json
```

### "props" option

Properties is accepting epsg.properties-like input and it would format it.

```shell
// will convert it to json
echo '1234=PROJCS[...' | ~/go/bin/parser props -in - -format json

// will pretty print it
~/go/bin/parser props -in epsg.properties -format pretty
```

There is the option to index entire content into own primitives like Datum, Authority, GeoCCS and etc. and print it as json.
Actual payload format is defined in [index_authorities.go:82](index_authorities.go).
Formatting option is `index`.

There is the option to parse only those definitions having `TOWGS84`.
Actual payload format is defined in [wgs84_repo.go:82](wgs84_repo.go).
Formatting option is `wgs84`.

### "template" option

Template option is executing golang template files with properties file as input.

The properties are formatted using specified visitor and resulting object passed to the template(s) as `repo` variable.

Format options are:

- `json` as provided by `to_map.go`
- `index` as provided by `index_authorities.go`
- `wgs84` as provided by `wgs84_repo.go`
- `ast` is antlr root node i.e. `wktcrsv1.IPropsFileContext`

For example, generating mappings for library [github.com/wroge/wgs84](https://github.com/wroge/wgs84) is performed as:
```shell
~/go/bin/parser template \
  -template templates/wgs84.tmpl \
  -in files/epsg.properties \
  -out out/epsg_db.go \
  -format wgs84 \
  -package wgs84
```
