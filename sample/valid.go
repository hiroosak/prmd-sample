package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/xeipuuv/gojsonschema"
)

func del(a []interface{}, i int) {
	s := a
	s = append(s[:i], s[i+1:])
	a = s
}

//
// schema.jsonのpropertiesを直接修正して,anyOfを1つに絞る
//
func getSchema() (map[string]interface{}, error) {
	target := flag.String("target", "user", "target schema")
	flag.Parse()

	schema := "../schema/schema.json"
	data, err := ioutil.ReadFile(schema)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}

	targets := fmt.Sprintf("#/definitions/%s", *target)

	var anyOf []interface{}
	anyOf = []interface{}{
		map[string]interface{}{"$ref": targets},
	}
	jsonData["properties"].(map[string]interface{})["body"].(map[string]interface{})["anyOf"] = anyOf

	fmt.Printf("properties : %s\n", targets)
	return jsonData, nil
}

func main() {
	var err error

	jsonData, err := getSchema()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	schemaDocument, err := gojsonschema.NewJsonSchemaDocument(jsonData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Loads the JSON to validate from a local file
	file := "./sample.json"
	jsonDocument, err := gojsonschema.GetFileJson("./" + file)
	fmt.Printf("Check ./%v\n", file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Try to validate the Json against the schema
	result := schemaDocument.Validate(jsonDocument)

	// Deal with result
	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		// Loop through errors
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}

}
