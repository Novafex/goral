package decl

// ExampleDeclaration is a base example template for a making new declaration
// files.
var ExampleDeclaration = Declaration{
	Name:        "My Resource",
	Description: "Enter your description of this resource type",
	Properties: []Property{
		{
			Name:        "Example",
			Description: "Enter a description of this property",
			Type:        PropertyTypeBoolean,
			Default:     true,
		},
	},
}
