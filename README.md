# Shopify Challenge Backend & SRE Summer 2021 Challenge

Submission for Shopify Challenge hosted at: [ec2-18-191-146-168.us-east-2.compute.amazonaws.com](https://ec2-18-191-146-168.us-east-2.compute.amazonaws.com). The goal of the challenge was to build a image repository, I chose to build a REST API with the following features:
#### Search for image(s)
  * By Text; in the form of tags
  * By ID; ID is automatically generated for every image that is added to the repo
  * By User; returns all public images of a given user
  * Default search; returns all public image
  
#### Add an image
  * I started by using **GridFS**, to store images and metadata, in **MongoDB**, but I did not like this implementation, as I wanted a remote URI so users can access images without being forced to download them. So, I used the same idea and split up the metadata and the actual image. All image metadata is stored in a **MongoDB** database. This metadata contains a key, that points to an **AWS S3 Bucket Link** that holds the actual image.
  * Users can tag images
  * **Computer Vision** is implemented to auto-tag all images, by making use of **Google's Vision API**
  * **Access Control** - Users can specify if any image is public/private
  
 #### Secure 
  * Users must create an account
  * Users are given a **JWT token** if they can **login**, that expires every **1 hour**
  * They must pass the given JWT Token with every request, to gain access to the data
  * User and passwords(**after hashing and salting**) are stored in MongoDB
  
 #### Deployment 
  * Deployed in a **Linux AWS EC2** instance
  * Build script lives in remote server, contains sensitive information, that is only exposed as environment variables at build time

## Technologies/Tools Used
* Golang (w/ mux for routing)
* AWS (EC2 & S3)
* MongoDB
* Google Vision API
* JWT Tokens


## Run the app locally
    
    # INIT ENVIRONMENT VARIABLES
    
    export AWS_ID=<TOKEN_HERE>
    export AWS_SECRET=<TOKEN_HERE>
    export SECRET_KEY=<TOKEN_HERE>
    export MONGO_USER=<TOKEN_HERE>
    export MONGO_PASSWORD=<TOKEN_HERE>
    export DB_NAME=<TOKEN_HERE>
    export GOOGLE_APPLICATION_CREDENTIALS=<PATH_TO_JSON_KEYS_FILE>
    export PORT=<PORT>
    
    # OPTION 1: RUN DIRECTLY
    
    go run main.go
    
    # OPTION 2: BUILD BINARY AND THEN EXECUTE
    
    go build
    ./shopify-challenge
    
  

# REST API 

Documentation to the given submission is given below

## Response

All responses include an appropriate status code. The format of the response is standardized. 

### Case: Success
 
    { 
      success:true, 
      data:<APPROPRIATE DATA HERE>
    }

### Case: Error
 
    { 
      success:false, 
      error:<APPROPRIATE ERROR MESSAGE HERE>
    }
    
## Supported Endpoints

**Note: Users are advised to use Postman to test the API** 

### `POST /signup`
   #### Body - `application/json`:
   ```
   {
     "username": "username length must be greater than 3",
     "password": "password length must be greater than 8"
   }
   ```
   #### Example Success Response
   ```
   {
     "success": true,
     "data": "YOUR_TOKEN_HERE"
   }
   ```
   The given credentials will be later on used to login
   
### `POST /login`
   #### Body - `application/json`:
   ```
   {
     "username": "a valid username",
     "password": "password matching given username"
   }
   ```
   #### Example Success Response
   ```
   {
     "success": true,
     "data": "JWT_TOKEN"
   }
   ```
   The given JWT token will be used in all further requests for authorization purposes.
   
### `POST /api/v1/add`
   #### Body - `multipart/formdata`:
   ```
     "img": <IMAGE_YOU_WANT_TO_ADD_TO_REPO>
     "private": true/false
     "tags": "commma,seperated,tags"
   ```
   #### Example Success Response
   ```
   {
     "success": true,
     "data": "IMAGE_ID"
   }
   ```
### `GET /api/v1/search`
   #### Body - `application/json':
   ```
   {
    "type": "by_id",
    "id": "IMAGE_ID"
   }
   
   OR 
   
   {
    "type": "by_tags",
    "tags": ["array", "of", "tags"]
   }
   
   OR 
   
   {
    "type": "by_user",
    "username": "a valid username"
   }
   
   OR
   
   {
    "type": "all"
   }
   ```
   #### Example Success Responses
   ```
   IF type == "by_id": 
   
   {
     "success": true,
     "data": {
      "key": "IMAGE_UUID",
      "tags": ["array", "of", "tags"],
      "owner": "username of owner",
      "private": true/false,
      "uri": "uri to s3 bucket containing image"
     }
   }
   
   ELSE: 
   
   {
     "success": true,
     "data": [
      {
       "key": "UUID for first search result",
       "tags": ["array", "of", "tags"],
       "owner": "username of owner",
       "private": true/false,
       "uri": "uri to s3 bucket containing image"
      },
      {
       "key": "UUID for second search result",
       "tags": ["array", "of", "tags"],
       "owner": "username of owner",
       "private": true/false,
       "uri": "uri to s3 bucket containing image"
      }
     ]

   }
   ```


