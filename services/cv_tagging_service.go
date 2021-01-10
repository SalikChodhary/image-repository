package services

import(
	"context"
	vision "cloud.google.com/go/vision/apiv1"
)

func getImageTags(s3URI string) ([]string, error) {
	var res []string

	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}
	image := vision.NewImageFromURI(s3URI)
	annotations, err := client.DetectLabels(ctx, image, nil, 10)

	if err != nil {
		return nil, err
	}

	for _, annotation := range annotations {
		res = append(res, annotation.Description)
	}

	return res, nil
}