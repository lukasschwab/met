// Package met provides a thin wrapper around the Metropolitan Museum of Art
// Collection API. View the full documentation: https://metmuseum.github.io
//
// All request Options here allow for the specification of a custom HTTP client:
//
//  cli := &http.Client{}
//  res, err := Search(ObjectsOptions{
//    HTTPOptions: HTTPOptions{Client: cli},
//    Q: "sunflowers"
//  })
package met
