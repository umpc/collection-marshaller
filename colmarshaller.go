package colmarshaller

import (
	"encoding/json"

	"github.com/tidwall/tile38/controller/collection"
	"github.com/tidwall/tile38/geojson"
)

type PackageCollection struct {
	*collection.Collection
}

type Package struct {
	Rows   []Row    `json:"rows"`
	Fields []string `json:"fields"`
}

type Row struct {
	Id     string    `json:"id"`
	Obj    []byte    `json:"obj"`
	Values []float64 `json:"values"`
}

func (c *PackageCollection) MarshalJSON() ([]byte, error) {

	pack := Package{
		Rows:   make([]Row, c.Count()),
		Fields: c.FieldArr(),
	}
	i := 0

	c.Scan(0, false,
		func(id string, obj geojson.Object, values []float64) bool {
			objBytes, _ := obj.MarshalJSON()
			pack.Rows[i] = Row{
				Id:     id,
				Obj:    objBytes,
				Values: values,
			}
			i++
			return true
		},
	)
    return json.Marshal(pack)
}

func (c *PackageCollection) UnmarshalJSON(b []byte) (err error) {

	pack := &Package{}
	if err = json.Unmarshal(b, pack); err != nil {
		return
	}

	for i := range pack.Rows {
		obj, err := geojson.ObjectAuto(pack.Rows[i].Obj)
		if err != nil {
			return err
		}
		c.ReplaceOrInsert(pack.Rows[i].Id, obj, pack.Fields, pack.Rows[i].Values)
	}
    return
}