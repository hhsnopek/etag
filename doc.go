// Package etag provides valid strong & weak etag generation.
//
// Example Usage
//
// The following is a complete example of a http.Handler
//    import (
//      "net/http
//      "encoding/json"
//
//      "github.com/hhsnopek/etag"
//    )
//
//    type Foo struct {
//      Bar string `json:"bar"`
//    }
//
//    func FooHandler(w http.ResponseWriter, r *http.Request) {
//      body, err := json.Marshal(&Foot{Bar: "baz"})
//      if err != nil {
//        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//        return
//      }
//
//      etag := etag.Generate(body, true)
//      w.Header().Set("Content-Type", "application/json")
//      w.Header().Set("Etag", etag)
//      w.WriteHeader(http.StatusOK)
//      w.Write(body)
//      return
//    }
package etag
