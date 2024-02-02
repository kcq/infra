// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xcX2/juBH/KgTbR2/kzaZ98Fuyl16Du94FmwAtsAgKWhpbvJVIHUk5MQx/94KkKFH/",
	"bNmxE6fwU2JzNOTM/GaGMyN4hUOeZpwBUxJPVjgjgqSgQJhP05wm0d1P+l/K8ARnRMV4hBlJAU/K1REW",
	"8GdOBUR4okQOIyzDGFKiH1PLTJNKJSib4/V6hIEtejnatd34USYVYSH0MvUIduMsCYum/KWXcbW+G18F",
	"aZYQ1X9ij2AXzmtNLDPOJBjrXY3H+k/ImQKm9L8kyxIaEkU5C/6QnOnvKn5/FTDDE/yXoIJEYFdlcCsE",
	"F3aPCGQoaKaZ4Am+IRHSRwSp8HqEr8afj7/nda5iYKrgisDS6c2vjr/5b1yhGc9ZpHf821uo+AHEAoQT",
	"c+0gYGx8yxZUcJYWu2eCZyAUtQAgCSUFFhoKtAuIz5CKAYHHZISpglR24GvkviBCkKX+7IWHOv+7SJtn",
	"RkG4LRIiFZJ5GIKUszxB5lE04wLN6QJY4witncugsW2fLXyyfJrQsM3o3zGoGESTBaIS2UcQF4izZImI",
	"EYFOE0DTpaFXQNJqrynnCRCGrTc61/1ehrYqaBZneVqPfCve6PW2KXdQtSE9phITPi8wNSN5ovDk+1Mr",
	"KhjzGsJdACUVUXkHXh/M993HA5anWsFGbM1X651ES71kPOZp1BGDN5tmRhmVMWgtGhmMjQyzlmFCHkH7",
	"wIYYmbURnnGREmUykfpyWWmUMgVzMJErBSnJvI8R3iZBsZHjoo97V2S9jhMnFJgahgRLeygwoZngKXqO",
	"aRhrz9KrLjmjUABR0Inbeobftp+jxkPNXrsflMrxdfgvUCQiykRwEkVU702S+5pW68dyTww61G/w3G+t",
	"wXp2om3z3tQTZlM6agnfrb8nK8CDvQu1zz90u4JBtVvznjRYfvfYVgB4/LUU33LGKJufHeegoDFBXSiI",
	"rlX7EI80BfQcA6uL9EwkKp7yw2dEFHxSNH2ta/tH8gxfINDauuMi1XONcurrh93oQMh5hSvVSpltJyiI",
	"u44w2JYFjz1MuavfO+oOH3GH6HWR/oAwqlV3vdDpDXungZhDmf2ELeJZ4fVpulcF6xF+dNbau8zy7P0u",
	"NdYmvA0pjEq7vqIq2hdLuyGls8hyBjxihfXelZGnrT3Kore2TbvK0oJCmAuqlg86mVnLXBtwPfIfwK5z",
	"FRtjAREg/uHSiYXff5UmwUV/xMDOkFXni5XKtJzXGf0Flo6Z6cDFQCJDWvTg/vPp+v7u0y+wrJ4m5inb",
	"oaFsxo3hqEr02u3lDbq+v8MjvAAhrcbGF58vxno7ngEjGcUT/OVifDHWsCQqNrIFwBbmnzl05NRfqVSI",
	"JIl/F9S40bg1Laa7CE/wz6BuNZdG++9yx95UCUaSJL/P8OT7lnaVV2asn1qQ7WhjlZEpWSIBKhcMorZ0",
	"VROxa/dSwkATVR24zbSayAeXEa4Fq+9Pa+0URPuwuUPawj/jssM0X00GQwQxeG6UXHXr3HNZmcf0SW94",
	"tGxYJs0TRTMiVKCvSJ9cFgMWcuO8Lj591Y+8KO/pR4fNyrQ8VKA+SSV0CK76hftdUrZUk81T1Tn9xMMf",
	"IIpEFBZU3jVwShkRyy6+kXlyRhPo46rXkFNh31X1axp1xkyhUMjTlLAIKY7gBcJc5zVWz3NkporU1xPx",
	"G3GupoyaDFWw5dM/IFSuW++39tct/708XG/Zd9a2az46EVHs3ddPzBPXIxstg5Up79bWsAmoLoiY7xFh",
	"G13TUmnnvC0KRn/81RP/KpLAlpn6qA3DXXWUR42KXhdG9vSRd2dLlh8l/n0DC5gtKnbR7/UKPofOc+g8",
	"h87DhM7AnFgGq+Iyvu69gf4Mqha2rKjm9ttzDTWebqoceVOVYfu4/WgroaslNGmjWmERvJRlvbYKZfPi",
	"9AmfIxUThWTM8yRCU6iuo89UxV0zUc3zzxyMxxUVgq5Zfp/NJCjsx4my9hu3Bz0dqWJ8DKDaEnfwHbwp",
	"747oLcbu22ivTgLpgavRu9PadRQhUuFkcz6rofxXW9IfG+kbc+Fm8DQyWEYfIBRdTn99f4ekXdvQ5Bja",
	"vGhE+mpfV/bvFeA/dyQnH9wkisCa8GhoXneiLwaS2L5CZ0D9p1lGYQzhj64Qatd7avnmFczgwPbYS9l3",
	"k9foOnBzkgGdCGGHJah6pEOIO2/xLXsSzRHeK/sSbVmNbsdDdDt+wzuA182qB8Xq4NsbGayav5m+fTsJ",
	"tiNh3c77RaVN9vQn48OjwkG2ru/bvvPVhpXFOKOjkPuYWKnFhGBVjVHXgYCZABnDhhz6zZLUZ7rwooBF",
	"xpuURIqmoMuHhC5gM7Luyr2/lTvvmmW9MfDh8meU2yN3FFnFihm92PFXTRXVtfMHZDqq0oUpoCSEnEVa",
	"vJS80DRP8eTL38fjEU4psx+7rpQd2XM9pCHRCHtWt1H9TYBjps49YSn9Kf2wVFU90pGqHrzFd0hV1fYH",
	"ylWVsB8s/lQHH9B0d3PkMlN5s6h2MKnb+Chpyr0I8MZZqrZtO0n5b2H8P+QoHyO1YBCsypcDdk5RTkk7",
	"ZqgSVA9u5/0TVPVmwzvmJ6eHk09P7i2NE8hOTUC6MDQgO1WkHVnp0Vt8y6xUvuTyymxUCXcqndDqRANn",
	"uhszSt0+H280selVoPNc4kPMJSpX/WhDCd8Va0EzWFVvCg0a7/Y7qaUo3fTRfwNpt9zsvbw0fNJb4u3k",
	"x7xD4qI36x0WFQ+p7nNwPQfXc3A9cHDdbQBcGnjz9LfD/V83B/ZjwbsOg/03tz/AJLj+pvNOY+BS0o84",
	"A94T+wcYCW+D/l7D4T3hf54Qn+KEuIZNg2uxcDjIRVK8nC4nQUAyegGX04sIFthD96r56xDSoKL+WxT1",
	"L/255cr7jQ3TbP5fAAAA//8MngF3+kMAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
