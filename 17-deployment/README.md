# Deployment

## Connect to EC2
```bash
ssh -i </directory/namafile.pem> <username-server>@<public-ipv4>

# example:
ssh -i ~/file.pem ubuntu@19.0.1.2
```

Jika ada error saat akses server via ssh (permission denied). coba ubah hak akses file.pem nya
```bash
chmod 400 file.pem
```

## Transfer file to Server using SCP
jika ingin melakukan transfer folder, tambahkan `-r`
```bash
scp -i </directory/namafile.pem> </directory/nama-file-trasfer> <username-server>@<public-ipv4>:/home/ubuntu

# example:
scp -i ~/file.pem map.go ubuntu@19.0.1.2:/home/ubuntu
```

## Update dan Upgrade Instance
```bash
sudo apt-get update
sudo apt-get upgrade
```

## Install mysql-client
```bash
sudo apt install mysql-client-core-8.0
```

## Connect to Mysql Server
Note: jangan lupa untuk buka port 3306 di pengaturan security inbound rule
```bash
mysql -h <db-endpoint> -P <port-db> -u <username> -p

# example:
mysql -h db-endpoint11221.amazonaws.com -P 3306 -u admin -p 
```

## Install docker on Ubuntu Server
```bash
sudo apt install docker.io
```

## Build Image
```bash
docker build -t <image-name>:<tag> .

# example:
docker build -t api-images:latest .
```

## Show Images List
```bash
docker images

docker images list
```

## Delete Images
```bash
docker rmi <image-id>

docker rmi <image-name>

# example:
docker rmi api-images
```

## Create Container from Image
```bash
docker run -d
-p <host-port>:<container-port>
-e <env-name>=<env-value>
-e <env-name>=<env-value>
-v <host-volume>:<container-volume>
--name <container-name> <image-name>:<tag>

# example:
docker run -p 80:8000 --name apiContainer api-images:latest
```

## Show Container
```bash
# melihat container yang sedang running
docker ps

# melihat seluruh container, termasuk yang sedang stop
docker ps -a
```

## Stop/Start Container
```bash
docker stop <container-name>

docker start <container-name>
```

## Remove Container
```bash
docker rm <container-name>

docker rm <container-id>

# example
docker rm apiContainer
```

## Docker Logs Container
melihat logs dari container. berguna untuk tracing ketika terjadi error di aplikasi/container.
```bash
docker logs <container-name>
```

## Login Docker Hub dan Push image ke Docker Hub
```bash
docker login -u <username>

docker build -t <username-dockerhub>/<image-name>:<tag> .

docker push <username-dockerhub>/<image-name>
```

## Pull Image dari container registry
```bash
docker pull <image-name>
```

## Jika terjadi error (permission denied) saat Run Docker di Ubuntu
```bash
sudo usermod -a -G docker ubuntu

or

sudo chmod 777 /var/run/docker.sock
```

# Notes CICD

## Clone repository 
- lakukan clone github repo di server. 
```bash
git clone <url-git-repo>

# example
git clone https://github.com/iffakhry/alta-be11-cicd.git
```

## Setup Github action 
- buat folder .github/workflows
- buat file `deploy.yml`
- push ke branch `main`
```bash
# sample deploy.yml
name: Deploy to EC2

on:
  push:
    branches:
      - main

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - name: executing connect to server using ssh key
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.HOST }}
          username: ${{ secrets.USERNAME }}
          key: ${{ secrets.KEY }}
          port: ${{ secrets.PORT }}
          script: |
            cd /home/ubuntu/alta-be11-cicd
            git pull origin main
            docker stop be11Api
            docker rm be11Api
            docker build -t be11-api:latest .
            docker run -d -p 8080:8000 -e SERVER_PORT=${{ secrets.SERVER_PORT }} -e DB_USERNAME=${{ secrets.DB_USERNAME }} -e DB_PASSWORD=${{ secrets.DB_PASSWORD }} -e DB_HOST=${{ secrets.DB_HOST }} -e DB_PORT=${{ secrets.DB_PORT }} -e DB_NAME=${{ secrets.DB_NAME }} --name be11Api be11-api:latest

```

## Konfigurasi SECRET Variable
```bash
HOST --> IP Public v4 Server
KEY --> isi ssh private key(gcp) atau file .pem(aws)
PORT --> 22
USERNAME --> ubuntu (aws) atau username (gcp)

DB_HOST --> (rds)db endpoint, (gcp)ip database server
DB_NAME --> database name
DB_USERNAME
DB_PASSWORD
DB_PORT --> 3306

SERVER_PORT 
```





