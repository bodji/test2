# Plik

Plik is an simple and powerful file uploading system written in golang.

### Main features
   - Multiple data backends : File, OpenStack Swift, WeedFS
   - Multiple metadata backends : File, MongoDB
   - Shorten backends : Recuce your uploads urls (is.gd && w000t.me available)
   - OneShot : Files are destructed after first download
   - Removable : Give the hability to uploader to remove files from upload
   - TTL : Option to set upload expiration
   - Password : Protect the upload with login/password (Auth Basic)
   - Comments : Add comments to upload (in Markdown format)
   - Yubikey : Protect the upload with your yubikey. You'll need an OTP per download

### Version
1.0-RC4


### Installation

##### From release
To run plik, it's very simple :
```sh
$ wget https://github.com/root-gg/plik/releases/download/1.0-RC4/plik-1.0-RC4.tar.gz
$ tar xvf plik-1.0-RC4.tar.gz
$ cd plik-1.0-RC4/server
$ ./plikd
```
Et voilà ! You have how a fully functionnal instance of plik ruuning on http://127.0.0.1:8080. You can edit server/plikd.cfg to adapt the params to your needs (ports, ssl, ttl, backends params,...)

##### From sources
For compiling plik from sources, you need a functionnal installation of Golang, and npm installed on your system.

First, get the project and libs via go get
```sh
$ go get github.com/root-gg/plik/server
$ cd $GOPATH/github.com/root-gg/plik/
```

As root user you need to install grunt, bower, and setup the golang crosscompilation environnement :
```sh
$ sudo -c "npm install -g bower grunt"
$ sudo -c "client/build.sh env"
```

And now, you just have to compile it
```sh
$ make build
$ make clients
```

### API
Plik server expose a RESTfull API to manage uploads and get files :

Managing uploads :
   - POST        /upload
   - GET         /upload/{uploadid}

Managing files :
   - POST        /upload/{uploadid}/file             (request must be multipart, with a part named "file" for file data)
   - GET/HEAD    /upload/{uploadid}/file/{fileid}    (head request just print headers, it does not count as a download (for oneShot uploads))
   - DELETE      /upload/{uploadid}/file/{fileid}    (only works if upload has "removable" option)

Nice links :
   - GET/HEAD    /file/{uploadid}/{fileid}/{filename}
   - GET         /file/{uploadid}/{fileid}/{filename}/yubikey/{yubikey}


Examples :
```sh
Create an upload (in the json response, you'll have upload id and upload token)
$ curl -X POST 127.0.0.1:8080/upload

Create a OneShot upload
$ curl -X POST -d '{ "OneShot" : true }' 127.0.0.1:8080/upload

Upload a file to upload
$ curl -X POST --header "X-UploadToken: M9PJftiApG1Kqr81gN3Fq1HJItPENMhl" -F "file=@test.txt" 127.0.0.1:8080/upload/IsrIPIsDskFpN12E/file

Get headers
$ curl -I 127.0.0.1:8080/file/IsrIPIsDskFpN12E/sFjIeokH23M35tN4/test.txt
HTTP/1.1 200 OK
Content-Disposition: filename=test.txt
Content-Length: 3486
Content-Type: text/plain; charset=utf-8
Date: Fri, 15 May 2015 09:16:20 GMT

```

### Cli client
Plik is shipped with a golang multiplatform cli client (downloadable in web interface) :
```sh
Simple upload
$ plik file.doc
Multiple files
$ plik file.doc project.doc
Archive and upload directory (using tar+gzip by default)
$ plik -a project/
Secure upload (OpenSSL with aes-256-cbc by deault)
$ plik -s file.doc

```

### Participate

You are free to implement other data/metadata/shorten backends and submit them via
pull requests. We will be happy to add them in the future releases.
