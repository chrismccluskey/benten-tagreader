package graphql

import (
	"context"
	"fmt"
	"log"

	"github.com/chrismccluskey/benten-tagreader/file"
	"github.com/chrismccluskey/benten-tagreader/models"
	"github.com/machinebox/graphql"
)

type GraphQLFileHandler struct {
	fileUsecase file.Usecase
}

func NewFileGraphQLHandler(fu file.Usecase) {
	/*handler := &GraphQLFileHandler{
		fileUsecase: fu,
	}*/
}

func (f *GraphQLFileHandler) Store(ctx context.Context, file models.File) error {
	//TODO: use context here
	// -- init
	client := graphql.NewClient("http://localhost:4000/graphql")

	// --- build request
	request := graphql.NewRequest(`
		mutation ($id: String!, $name: String!, $size: Int!) {
			addFile(id: $id, name: $name, size: $size) {
				id
				name
				size
			}
		}
	`)
	request.Var("id", file.ID)
	request.Var("name", file.Name)
	request.Var("size", file.Size)
	request.Header.Set("Cache-Control", "no-cache")

	// --- send request
	var respData map[string]interface{}
	if err := client.Run(ctx, request, &respData); err != nil {
		log.Fatal(err)
	}

	// --- handle response
	fmt.Println(respData)

	return nil
}
