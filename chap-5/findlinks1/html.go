package main

//
// import "io"
//
// type Attribute struct {
// 	Key, Val string
// }
//
// type NodeType int32
//
// type Node struct {
// 	Type                    NodeType
// 	Data                    string
// 	Attr                    []Attribute
// 	FirstChild, NextSibline *Node
// }
//
// const (
// 	ErrorNode NodeType = iota
// 	TextNode
// 	DocumentNode
// 	ElementNode
// 	CommentNode
// 	DoctypeNode
// )
//
// func Parse(r io.Reader) (*Node, error) {
// 	return new(Node), nil
// }