#  File store application

### http server is used as store and client utility is used to update the files on the server

### Design Diagram
<img width="973" alt="file-store" src="https://user-images.githubusercontent.com/45117013/205504584-6ed87fcf-a180-446c-aa62-8d0eb695ff4f.png">

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
* Run the container using bind mount, below command will mount the store-files  directory inside the container to the store-files inside current working directory. Thus user can access the server-files  where files are stored and served. 
```
mkdir store-files
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
   
#### Word frequency
---
  Get most frequent or least frequent words.
* **URL:**
  /store/frequency
* **Method:**
  `GET`
  
*  **URL Params**
    None
* **Data Params**
   **Required:**
    `{
    "order": "dsc|asc",
    "limit": <number>
    }`
   
# How to run the client application
```
git clone https://github.com/aanand01762/file-store.git
cd file-store/client
cd config
```
* Update the server hostname and port number in the config.yaml or create new file, default values are localhost and 8080 respectively.
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
Using config file: ../file-store/client/config/config.yaml
[{"filename":"a","msg":"Files uploaded successfully: a"},{"filename":"a.py","msg":"Files uploaded successfully: a.py"},{"filename":"aman_k8s_config.yaml","msg":"Files uploaded successfully: aman_k8s_config.yaml"}]
```
## Remove files using client
```
% ./store rm a aman_k8s_config.yaml b 
Using config file: ../file-store/client/config/config.yaml
a removed successfully

aman_k8s_config.yaml removed successfully

b removed successfully
```
## List files using client
```
akumar32@akumar32XMD6M client % ./store ls                                   
Using config file: ../file-store/client/config/config.yaml
a a.py aman_k8s_config.yaml 
```
## Get count of words in all files using client
```
% ./store wc
Using config file: ../file-store/client/config/config.yaml
228
```
##  Update the file on server using client
* If file exist with same name and different content it will updated with uploaded file
```
client % ./store update  ~/personal-git/test/a.py 
Using config file: ../file-store/client/config/config.yaml
Updated the content of the file 'a.py' with latest value
```
* If file exists with same content type but different name it will be replaced with uploaded file
```
client % ./store update  ~/personal-git/test/b   
Using config file: ../file-store/client/config/config.yaml
Changed file name from: 'a.py' to new file name: 'b' because both files had same content
```

##  List most frequent or least frequent words in all files on the server. 
* Pass --limit flag for number of words, --order flag with asc value for least frequent and dsc for most frequent words. Default values for limit  and order are 10 and "dsc" respectively.
```
akumar32@akumar32XMD6M client % ./store freq-words --order asc --limit 15  
Using config file: ../file-store/client/config/config.yaml
waitforqueuestodraintask always headersheaders server initedgeclientcachetask validatesequentialkeystask response verifyfalse latestconfendtask datapayload printresponsetext initservercachetask initserverregiontask authorization config 
```

# How to run the application inside kubernetes cluster
* Update  the hostPath field in the deployment.yaml(inside server folder) to mount the the host 
```
    hostPath:
          path: <host path>
          type: DirectoryOrCreate
```
* Apply to run the application as container inside k8s pod.
```
cd sever
kubectl apply -f deployment.yaml
```
* Now create the service of type NodePort. By default service will use 30000 port of the host node, but user can always update and use with different values.
```
kubectl apply -f serive.yaml
```
