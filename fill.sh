#!/bin/bash
HOST=localhost

curl -X POST ${HOST}:3000/languages -d '{"language":"German","level":"B2","description":"school"}'
curl -X POST ${HOST}:3000/languages -d '{"language":"English","level":"B2","description":"school"}'

curl -X POST ${HOST}:3000/education -d '{"school":"pwr","degree":"engineer","field":"cybersec"}'
curl -X POST ${HOST}:3000/education -d '{"school":"pwr","degree":"masters","field":"cybersec"}'


curl -X POST ${HOST}:3000/experience -d '{"company":"Nokia","role":"Linux Admin"}'
curl -X POST ${HOST}:3000/experience -d '{"company":"Docusoft","role":"student"}'

curl -X POST ${HOST}:3000/projects -d '{"project_name":"k8s automation","technology_used":"ansible","description":"wooooooooooot"}'
curl -X POST ${HOST}:3000/projects -d '{"project_name":"cv api","technology_used":"Go","description":"literally this"}'