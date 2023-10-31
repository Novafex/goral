// This is my super cool testing struct for Goral development
// 
// 
// © 2023 My Name
// 
// Some dummy license
package goral

// This is my super cool testing struct for Goral development
type Test struct {
    // Cool name to give it
    Name string `json:"name" sql:"name"`

    // Test property for optionals
    Foo *string `sql:"foo" json:"foo"`

    // This one is to test the casing and such
    SomethingMoreRealistic *float32 `json:"somethingMoreRealistic" sql:"something_more_realistic"`
}
