package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
	"github.com/weaviate/weaviate/entities/models"

	"github.com/rchang0501/media-search-engine/internal"
)

func addModelToSchema(client *weaviate.Client, classObj *models.Class) {
	// add the schema
	err := client.Schema().ClassCreator().WithClass(classObj).Do(context.Background())
	if err != nil {
		panic(err)
	}
}

func addObjectsToModel(client *weaviate.Client) {
	objs := []*models.Object{}

	for _, imgPath := range internal.Images {
		base64String := internal.ConvertImageToBase64(imgPath)

		imgName := internal.ExtractImageName(imgPath)

		// Create a new object
		obj := &models.Object{
			Class: "Picture",
			Properties: map[string]interface{}{
				"image": base64String,
				"text":  fmt.Sprintf("This is a picture of a %s", imgName),
			},
		}

		objs = append(objs, obj)
	}

	// batch write items
	batchRes, err := client.Batch().ObjectsBatcher().WithObjects(objs...).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, res := range batchRes {
		if res.Result.Errors != nil {
			panic(res.Result.Errors.Error)
		}
	}
}

func retrieveData(client *weaviate.Client, imagePath string) {
	// Convert image data to base64
	base64Image := internal.ConvertImageToBase64(imagePath)

	fields := []graphql.Field{
		{Name: "image"},
		{Name: "text"},
	}

	nearText := client.GraphQL().
		NearImageArgBuilder().WithImage(base64Image)

	// Perform GraphQL query
	resImage, err := client.GraphQL().Get().
		WithClassName("Picture").
		WithFields(fields...).
		WithNearImage(nearText).
		WithLimit(1).
		Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	resStr := fmt.Sprintf("%+v", resImage)
	resultImage := internal.ExtractImageFromJSON(resStr)

	fmt.Println("Result image:", resultImage)

	resultData, err := base64.StdEncoding.DecodeString(resultImage)
	if err != nil {
		log.Fatal(err)
	}

	resultImagePath := "imgs/result.jpg"
	err = os.WriteFile(resultImagePath, resultData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result image written to", resultImagePath)
}

func main() {
	cfg := weaviate.Config{
		Host:   "localhost:8080",
		Scheme: "http",
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	schema, err := client.Schema().Getter().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("schema: %+v", schema)
}
