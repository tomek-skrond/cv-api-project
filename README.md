# CV API

This is a Go API that is my CV as an API.

It is currently under development, for now it contains the following endpoints:
- `/education` -> endpoint with my Education info (GET and POST implemented)


I am developing code in `master` branch, in `dockerize` branch, `dockerize` branch contains `docker-compose` configuration for deploying this API.


## `Education` table data format
Below you can find the data format for table `education`.

To retrieve the data use a command:
```
curl -ks localhost:3000/education
```

Data:
```
[
  {
    "ID": 1,
    "school": "Politechnika Wroclawska",
    "degree": "Masters Degree",
    "field": "",
    "dateStarted": "2021-02-18T00:00:00Z",
    "dateEnded": "2023-02-18T00:00:00Z"
  },
  {
    "ID": 2,
    "school": "Politechnika Wroclawska",
    "degree": "Engineers Degree",
    "field": "",
    "dateStarted": "2021-02-18T00:00:00Z",
    "dateEnded": "2023-02-18T00:00:00Z"
  }
]
```
To insert another row into table, use the following command:
```
curl -X POST -d '{"ID":1,"school":"Politechnika Wroclawska","degree":"Engineers Degree","dateStarted":"2021-02-18T00:00:00Z","dateEnded":"2023-02-18T00:00:00Z"}' http://localhost:3000/education
```
