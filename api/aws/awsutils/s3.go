package awsutils

import (
	"context"
	"log"
	"os"
	"time"

	//	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (t *AWSUtils) S3Download(bucket, item, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Unable to open file %q, %v", err)
		log.Fatal("Open Faile Failed")
	}

	// 要調整
	download := func(ctx context.Context) {
		downloader := s3manager.NewDownloader(t.session)
		ctx2 := context.WithTimeout(ctx)
		_, err := downloader.DownloadWithContext(
			ctx,
			file,
			&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(item),
			},
		)
		if err != nil {
			log.Printf("Unable to download item %q, %v", item, err)
		}
		log.Println("ダウンロード完了")
		ctx.Done()
	}

	log.Println("Start download")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	go download(ctx)

	select {
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			log.Printf("ERROR: ", err)
		}
		log.Printf("Done")
	}

	// log.Println("Downloaded", file.Name(), numBytes, "bytes")
	//	for {
	//		select {
	//		case <-ctx.Done():
	//			log.Println("ダウンロード完了")
	//		default:
	//			fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	//		}
	//		time.Sleep(time.Second)
	//	}

}
