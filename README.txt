minikube stop
minikube delete
minikube start --cpus 4 --memory 8192

~~~~~~~~~~~~~~~~~~~~~~~`~
POSTGRES

creazione di un postgres con singolo pod: manca il PVC che invece troviamo qui: 
(https://github.com/shri-kanth/kuberenetes-demo-manifests/blob/master/mysql-deployment.yaml)
https://medium.com/@xcoulon/deploying-your-first-web-app-on-minikube-6e98d2884b3a

in alternativa si puo' usare il postgres operator di zalando:
https://postgres-operator.readthedocs.io/en/latest/quickstart/


~~~~~~~~~~~~~~~~~~~~~~~~~~
MONGO

docker pull mongo
# docker run -d -p 27017-27019:27017-27019 --name mongodb mongo:4.0.20
docker run -d -p 27017:27017 --name mongodb mongo
docker exec -it mongodb bash
root@fe91c424d518:/# mongo
> show dbs
> use people_db
> db.people.save({ firstname: "Nic", lastname: "Raboy" })
> db.people.save({ firstname: "Maria", lastname: "Raboy" })
> db.people.save({ firstname: "Maria", lastname: "Raboy" })
-> { "_id" : ObjectId("5fa82cd7cf5ad7f2ab6ccdf6"), "firstname" : "Nic", "lastname" : "Raboy" }
   { "_id" : ObjectId("5fa82ce1cf5ad7f2ab6ccdf7"), "firstname" : "Maria", "lastname" : "Raboy" }

> db.people.find({ firstname: "Nic" })
-> { "_id" : ObjectId("5fa82cd7cf5ad7f2ab6ccdf6"), "firstname" : "Nic", "lastname" : "Raboy" }   


mkdir mongo_example && cd mongo_example
go mod init podcast
touch main.go
mkdir environment handler repository dto model
ln -s "/Applications/Sublime Text.app/Contents/SharedSupport/bin/subl" /usr/local/bin/subl
