#  File store application

## http server is used as store and client utility is used to update the files on the server

#### How to run the server application
```
git clone https://github.com/aanand01762/file-store.git
cd file-store/server
mkdir store-files 
go build
./url-shortner
```

#### How to run the application inside container
* Build container image using Dockerfile
```
git clone  https://github.com/aanand01762/file-store.git
cd file-store/server
mkdir store-files 
docker build .
```
* Tag the image
```
url-shortner % docker images
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
<none>       <none>    0e67cb6971c4   16 hours ago   310MB
url-shortner % docker tag  0e67cb6971c4 file-store:v1
url-shortner % docker images
REPOSITORY     TAG       IMAGE ID       CREATED        SIZE
file-store  v1        0e67cb6971c4   16 hours ago   310MB
```
* Run the container using bind mount, below command will mount the store-files  directory inside the container to the current working directory. Thus user can access the server-files  where files are stored and served. 
```
docker run -d -it --name <container_name> -p <localhost port>:8080 --mount type=bind,source="$(pwd)",target=/app/store-files  file-store:v1
```


## API Contracts

#### Add File
---
  Add single or multiple files to http server, Return error message if file name or file with same content already exists on server.
* **URL:**
  /store/add
* **Method:**
  `POST`
*  **URL Params**
   None
* **Body**
   **Required:**
    `body type: form data`
    ***Format***
     `key : multiplefiles`
     `value : <files>`    
  
#### Remove file
---
  Delete file from the store.
* **URL:**
  /store/delete
* **Method:**
  `DELETE`
  
*  **URL Params**
    None
* **Data Params**
   **Required:**
    `{
        filename: string
      }`
  
  
#### List Files
---
  Returns list of files on the server.
* **URL:**
  /store/files
* **Method:**
  `GET`
* **Data Params**
  None
  
#### Update file
---
  Update the file with new file, if filename or another file with same file content exists then it is replaced with latest file with appropriate message.
* **URL:**
  /store/update
* **Method:**
  `PUT`
* **Data Params**
   **Required:**
    `body type: form data`
    ***Format***
     `key : file`
     `value : <file>`
     
#### Word Count
---
  Returns the total number of words in all files.
* **URL:**
  /store/word-count
* **Method:**
  `GET`
* **Data Params**
   NONE
   
# How to run the client application
```
git clone https://github.com/aanand01762/file-store.git
cd file-store/client
cd config
```
* Update the server hostname and port number, default value is localhost and 8080 respectively.
```
HOST: <hostname>
PORT: <port number>
```
* Now build the CLI. Note: Go and Cobra  and viper package should be present in go path
```
go get github.com/spf13/cobra
go get github.com/spf13/viper
go build
./store 
```

### Sample output
```
akumar32@akumar32XMD6M client % ./store   
client app to manage file on a http server

Usage:
   [command]

Available Commands:
  add         store add <file1> <file2> 
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ls          store ls
  rm          store rm file.txt
  update      store update file.txt
  wc          store wc

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```
## Add files using client 
```
% ./store add ~/a ~/a.py ~/aman_k8s_config.yaml 
Using config file: /Users/akumar32/personal-git/file-store/client/config/config.yaml
[{"filename":"a","msg":"Files uploaded successfully: a"},{"filename":"a.py","msg":"Files uploaded successfully: a.py"},{"filename":"aman_k8s_config.yaml","msg":"Files uploaded successfully: aman_k8s_config.yaml"}]
```
## Remove files using client
```
% ./store rm a aman_k8s_config.yaml b 
Using config file: /Users/akumar32/personal-git/file-store/client/config/config.yaml
a removed successfully

aman_k8s_config.yaml removed successfully

b removed successfully
```

