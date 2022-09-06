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


