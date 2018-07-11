# goShaula
- golangでSocket通信の実装


- golangでPortScanの実装
```
$ go run portScan.go
scanning port 20-30000...
 22 [open]  -->   SSH
 25 [open]  -->   SMTP
 53 [open]  -->   <unknown>
 80 [open]  -->   web service nginx
 88 [open] -->  kerberos
 139 [open]  -->   netbios
 445 [open]  -->   Samba
 548 [open]  -->   <unknown>
 587 [open]  -->   <unknown>
 631 [open]  -->   cups
 2181 [open]  -->   <unknown>
 5800 [open]  -->   VNC remote desktop
 27017 [open]  -->   mongodb [ http://www.mongodb.org/ ]
```