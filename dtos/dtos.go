// dtos/dtos.go
package dtos

type Post struct {
	PostID          int32
	Title           string
	Content         string
	Author          string
	PublicationDate string
	Tags            []string `json:"tags"`
}
