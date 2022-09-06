# Deployment

## Connect to EC2
```bash
ssh -i </directory/namafile.pem> <username-server>@<public-ipv4>

example:
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

example
scp -i ~/file.pem map.go ubuntu@19.0.1.2:/home/ubuntu
```