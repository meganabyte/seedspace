package queries

import (
	"github.com/graphql-go/graphql"
	"api/queries/fields"
)

// RootQuery is the entry point for graphql queries
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
	  "project": fields.GetProject,
	  "allProjects": fields.GetAllProjects,
	  "student": fields.GetStudent,
	  "guest": fields.GetGuest,
	},
})


