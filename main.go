package main

import (
    "bytes"
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "net/http"
		"time"
		"image"
	//	"encoding/json"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/gridfs"
		"go.mongodb.org/mongo-driver/mongo/options"
		"github.com/gorilla/mux"
		"github.com/SalikChodhary/shopify-challenge/services"
)

func InitiateMongoClient() *mongo.Client {
    var err error
    var client *mongo.Client
    uri := "mongodb+srv://admin:dbtest@cluster0.xplun.mongodb.net/ShopifyChallenge?retryWrites=true&w=majority"
    opts := options.Client()
    opts.ApplyURI(uri)
    if client, err = mongo.Connect(context.Background(), opts); err != nil {
        fmt.Println(err.Error())
    }
    return client
}
func UploadFile(file, filename string) {

    data, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
    conn := InitiateMongoClient()
    bucket, err := gridfs.NewBucket(
        conn.Database("myfiles"),
    )
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    uploadStream, err := bucket.OpenUploadStream(
        filename,
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer uploadStream.Close()

    fileSize, err := uploadStream.Write(data)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
}
func DownloadFile(fileName string) {
    conn := InitiateMongoClient()

    // For CRUD operations, here is an example
    db := conn.Database("myfiles")
    fsFiles := db.Collection("fs.files")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    var results bson.M
    err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
    if err != nil {
        log.Fatal(err)
    }
    // you can print out the results
    fmt.Println(results)

    bucket, _ := gridfs.NewBucket(
        db,
    )
    var buf bytes.Buffer
    dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("File size to download: %v\n", dStream)
    ioutil.WriteFile(fileName, buf.Bytes(), 0600)

}

// func uploadFileToS3(s *session.Session, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
// 	size := fileHeader.Size
//   buffer := make([]byte, size)
// 	file.Read(buffer)
// 	println(http.DetectContentType(buffer))
	
// 	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
// 		Bucket:               aws.String("salik-test-bucket"),
// 		Key:                  aws.String(fileHeader.Filename),
// 		ACL:                  aws.String("public-read"),// could be private if you want it to be access by only authorized users
// 		Body:                 bytes.NewReader(buffer),
// 		ContentLength:        aws.Int64(int64(size)),
// 		 ContentType:        	aws.String(http.DetectContentType(buffer)),
// 		ServerSideEncryption: aws.String("AES256"),
// 		StorageClass:         aws.String("INTELLIGENT_TIERING"),
//  })
//  if err != nil {
// 	return "", err
// }

// return fileHeader.Filename, err
// }

func tempPost(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(int64(2048000))

	file, fileHeader, err := r.FormFile("img")

  if err != nil{
    log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
    return
	}
	
	_, _, err = image.Decode(file)
  if err != nil {
    log.Printf("could not decode body into an image")
    w.Header().Add("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("could not decode body image"))
    return
	}

	err = services.InitS3Instance()

	if err != nil {
		fmt.Fprintf(w, "Could not upload file")
	}

	err = services.UploadFileToS3(file, fileHeader)

	if err != nil {
		fmt.Fprintf(w, "Could not upload file")
	}

	fmt.Fprintf(w, "Image uploaded successfully.")

}

func main() {
		router := mux.NewRouter()
		router.HandleFunc("/api/post", tempPost).Methods("POST")
		log.Fatal(http.ListenAndServe(":8000", router))
}