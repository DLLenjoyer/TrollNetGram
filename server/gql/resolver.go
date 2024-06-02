package gql

import (
	"github.com/DLLenjoyer/TrollNetGram/server/models"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Resolver struct {
	db *gorm.DB
}

func NewResolver(db *gorm.DB) *Resolver {
	return &Resolver {
		db: db,
	}
}

func (r *Resolver) RootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": r.usersQuery(),
		},
	})
}

func (r *Resolver) RootMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"registerUser": r.registerUserMutation(),
		},
	})
}

func (r *Resolver) usersQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(userType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var users []models.User
			if err := r.db.Find(&users).Error; err != nil {
				return nil, err
			}
			return users, nil
		},
	}
}


func (r *Resolver) registerUserMutation() *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			name := params.Args["name"].(string)
			email := params.Args["email"].(string)
			password := params.Args["password"].(string)

			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return nil, err
			}

			user := models.User {
				Name: name,
				Email: email,
				Password: string(hashedPassword),
			}
			if err := r.db.Create(&user).Error; err != nil {
				return nil, err
			}

			return user, nil

		},
	}
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.ID},
		"name": &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
	},
})
