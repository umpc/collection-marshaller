package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "strconv"
    "time"

    "github.com/tidwall/tile38/controller/collection"
    "github.com/tidwall/tile38/geojson"

    cmarshaller "github.com/umpc/collection-marshaller"
)

const MB = 1000000

func main() {
    geoindex := collection.New()

    // Wrap the collection in a struct with marshalling methods
    pack := cmarshaller.PackageCollection{ geoindex }

    // Populate the collection
    fmt.Println("Inserting features...")
    insertFeatures(geoindex)

    fmt.Printf("Objects in geoindex: %v\n",
        geoindex.Count())

    // Marshal from pack into []byte slice
    b, err := pack.MarshalJSON()
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Marshaled pack. Size of b: %v MB\n",
        float64(len( b )) / MB)

    // Change to a new index, overwriting the pointer for the simplicity of this example.
    geoindex = collection.New()

    // Update the type wrapper
    pack = cmarshaller.PackageCollection{ geoindex }

    fmt.Printf("Created new geoindex. Objects in geoindex: %v\n",
        geoindex.Count())

    // Unmarshal into pack from []byte slice
    if err = pack.UnmarshalJSON(b); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Unmarshaled from b. Objects in geoindex: %v\n",
        geoindex.Count())
}

// Inserts points ~14.731km apart.
func insertFeatures(geoindex *collection.Collection) {
    randSrc := rand.NewSource(time.Now().UnixNano())

    for lng := float64(-180); lng < 180; lng += 0.25 {
        for lat := float64(-90); lat < 90; lat += 0.25 {

            id := strconv.FormatFloat(lng, 'f', -1, 64) + strconv.FormatFloat(lat, 'f', -1, 64)
            objBytes, _ := json.Marshal(map[string]interface{}{
                "type": "Feature",
                "geometry": map[string]interface{}{
                    "type": "Point",
                    "coordinates": []interface{}{lng,lat},
                },
                "properties": map[string]interface{}{
                    "randomNumber": rand.New(randSrc).Float64(),
                },
            })
            obj, _ := geojson.ObjectAuto(objBytes)

            geoindex.ReplaceOrInsert(id, obj, nil, nil)
        }
    }
}