# Marshaller methods for Tile38's collection package

This is just a simple wrapper for saving and restoring a collection to and from JSON, without modifying the collection API.

```
geoindex := collection.New()

pack := colmarshaller.PackageCollection{ geoindex }

b, err := pack.MarshalJSON()
if err != nil {
    fmt.Println(err)
    return
}

if err = pack.UnmarshalJSON(b); err != nil {
    fmt.Println(err)
    return
}
```
