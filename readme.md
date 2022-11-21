# Cake Store RestAPI

dev with go to save cake data


## How To Run

1. To run this project, make sure you have install docker.
then you can run command below

```bash
  docker run --name maria-cakestore -e MARIADB_ROOT_PASSWORD=root -e MARIADB_DATABASE=cake_store -p 3306:3306 -d mariadb/server:10.3
```
2. After that on console db , create table with SQL below
```bash
    create table if not exists cake(
        id int auto_increment primary key ,
        title varchar(255) not null ,
        description text ,
        rating int ,
        image text,
        is_active boolean default 1,
        created_at timestamp default now(),
        updated_at timestamp default now()
    );
```
3. clone and open project and run
```bash
  go run main.go
```
4. and run in postman with port 9000 ex : 
```bash
  localhost:9000
```





## API Reference

#### Get List Cake

```http
  GET /cakes
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `limit` | `string` | **Optional**. count limit data |
| `page` | `string` | **Optional**. show what page |

#### Get Detail Cake

```http
  GET /cakes/{id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of cake to fetch |

#### Add New Cake
```http
  POST /cakes
```
in body raw add payload json like below
```json
{
    "title" : "cake 3",
    "description": "desc cake",
    "rating" :7,
    "image" : "link image"
}
```
#### Edit  Cake
```http
  PUT /cakes/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of cake to edit |

in body raw add payload json like below
```json
{
    "title" : "cake 3",
    "description": "desc cake",
    "rating" :7,
    "image" : "link image"
}
```

#### Delete Cake
```http
  DELETE /cakes/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of cake to delete |

