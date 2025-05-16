# API Growth 💙 🐿️ 🐍 🦀
This repository was created to make projects available in various programming languages ​​for educational purposes and to collaborate with the developer community. A joke that was born on social media and materialized in this repository ❤️.

Programming languages ​​❤️ are tools and should be used to solve specific problems than what they were proposed to solve. But we know that it goes far beyond this 😍, in this equation we have to add a pinch of LOVE 😍 and when you have this combination things start to get even more interesting 😂😂.

---
The scope of the project is to create a rEST API, a CRUD and persist it in memory and place it in a docker image. The size of this docker image could not exceed 6Mb, but we are aware of the limitations that each language has and in this regard you can send a larger image, try to make it as small as you can and very lean ☺️.

Your POST will receive a JSON of 1mb or 3mb and persist it in memory.
Below is an example and description of what you will need to implement in the API.

The entire repo was organized by programming languages, feel free to collaborate by sending us a pull request, below we will leave the documentation on how to make a PR.

What we will send to [POST] will be a 1Mb or 3Mb json with more than 40k lines and the body of the json is below:

```bash
[
   {
      "Country":"BRZ",
      "Indicator":"NGDP_R",
      "Value":183.26,
      "Year":2002
   },
   {
      "Country":"AFG",
      "Indicator":"NGDP_R",
      "Value":198.736,
      "Year":2003
   }
]
```
## Pull Request

You can organize your directory like the examples below:
```bash
grow.go/
└── jeffotoni
    ├── grow.fiber
    │   └── README.md
    └── grow.standard.libray
        ├── Dockerfile
        ├── go.mod
        ├── main.go
        ├── main_test.go
        └── README.md

```
You can organize your project by choosing the language you will implement and then your github user and within your directory you can create and organize your contributions.

Check out more examples:
```bash
grow.python/
└── cassiobotaro
    ├── Dockerfile
    ├── main.py
    ├── README.md
    └── requirements.txt
```

```bash
grow.rust
└── marioidival
    └── actix
        ├── Cargo.toml
        └── src
            └── main.rs
```
## Docker
You can use Docker or Podman to create your images, remembering that the smaller the better, so try to make the smallest images possible.
We will execute the following command:
```bash
$ docker build --no-cache -f Dockerfile -t growth/<lang>:latest .
```
And then we will run it:
```bash
$ docker run --rm -it -p 8080:8080 growth/<lang>
```

Feel free to play with the possibilities, you can use docker-compose too, you can use the scale option if you want space for creativity is always welcome 😁.

## Tests Stress
We will be stress testing your project, so be sure to take this into consideration. We will be using V6 and Locust for the tests and they are located in the root of the repository with the installation and configuration manual. With our example ready and beautiful, just run it 😍.

## Endpoints to be implemented
The endpoints that must be implemented are listed below, we will follow the same pattern for all projects:

#### POST
Creating our database in memory, this request is asynchronous and will run in the background, but only implement this feature if your language provides support.
```bash
$ curl -i -XPOST -H "Content-Type:application/json" \
localhost:8080/api/v1/growth -d @3mb-growth_json.json
{"msg":"In progress"}
```

#### GET
With this endpoint we can view the status of the processing we sent in [POST]
```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/post/status
{"msg":"complete","test value"":183.26, "count":42450}
```

#### GET
This endpoint searches memory to return the result.
```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002
{"Country":"BRZ","Indicator":"NGDP_R","Value":183.26,"Year":2002}
```

#### PUT
This endpoint will update the database in memory, if the data does not exist it will create a new one.

```bash
$ curl -i -XPUT -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002 \
-d '{"value":333.98}'
```
#### GET
Making a request to check if what we changed or created new is in the database.

```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002
{"Country":"BRZ","Indicator":"NGDP_R","Value":333.98,"Year":2002}
```
#### DELETE
This endpoint will remove the data from our memory database.

```bash
$ curl -i -XDELETE -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/brz/ngdp_r/2002 
```

#### GET
This endpoint will return the size of our database in memory.

```bash
$ curl -i -XGET -H "Content-Type:application/json" \
localhost:8080/api/v1/growth/size
{"size":42450}
```