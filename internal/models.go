package internal

import "github.com/weaviate/weaviate/entities/models"

var ImageModel = &models.Class{
	Class:           "Picture",
	Vectorizer:      "img2vec-neural",
	VectorIndexType: "hnsw",
	ModuleConfig: map[string]interface{}{
		"img2vec-neural": map[string]interface{}{
			"imageFields": []string{"image"},
		},
	},
	Properties: []*models.Property{
		{
			Name:     "image",
			DataType: []string{"blob"},
		},
		{
			Name:     "text",
			DataType: []string{"string"},
		},
	},
}
