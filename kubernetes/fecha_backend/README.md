Sistema que tiene un servicio front y otro bakend

backend - api rest en go


HOW-TO
------

- Descargar imagen docker de go
docker pull golang

- Crear api rest: 
jmoreno25/kubernetes/fecha_backend/backend/src/main.go

- Comprobar docker
docker run --rm -dti -v $PWD/:/go --net host --name go golang bash
    -v $PWD/:/go --> añadimos al contenedor el dorectorio go con el contenido de donde estemos

docker ps -l

docker exec -ti 6c0d1016237ae5858165eef453e3293662665ad687bd9b8dc0d5bd5162843c09

root@KM-5CD242HTHV:/go# go run main.go