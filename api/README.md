# Registry

# Endpoints

- [<code>POST</code> registry](#post-registry)
- [<code>GET</code> registry](#get-registry)
- [<code>GET</code> all registries](#get-all-registry)
- [<code>DELETE</code> registry](#delete-registry)
- [<code>PUT</code> registry](#put-registry)

## Examples
# Post registry
- **Request**
``` json
{
	"data":{
		"nombre":"jun carlos", 
		"apellido":"jun carlos", 
		"correo":"juanki@mail", 
		"telefono": "783278234", 
		"type_customer":1
		}
}
```
- **Response**
``` json
{
    "id": 1,
    "nombre": "jun carlos",
    "apellido": "jun carlos",
    "correo": "estilista@gmail.com",
    "telefono": "783278234",
    "type_customer": 1,
    "created_at": "2019-03-23 09:20:07"
}
``` 
# Get all registry
- **Request**
- <code>localhost:8080/registry</code>
- **Response**
``` json
[
    {
        "id": 1,
        "nombre": "jun carlos",
        "apellido": "jun carlos",
        "correo": "estilista@gmail.com",
        "telefono": "783278234",
        "type_customer": 1,
        "created_at": "2019-03-23 09:20:07"
    },
    {
        "id": 2,
        "nombre": "Ricardo Tapia",
        "correo": "estilista@gmail.com",
        "telefono": "78687698",
        "type_customer": 1,
        "created_at": "2019-03-23 09:20:13"
    }
```

# Get registry
- **Request**
- <code>localhost:8080/registry/1</code>
- **Response**
``` json
{
    "data": {
        "id": 5,
        "nombre": "Elver",
        "apellido": "Galarga",
        "correo": "elver@galarga.com",
        "telefono": "1234567890",
        "type_customer": 0,
        "created_at": "2019-04-27 09:58:00"
    }
}
```

# Put registry
- **Request**
- <code>localhost:8080/registry/1</code>
``` json
{
	"data":{
		"nombre":"jun carlos", 
		"apellido":"jun carlos", 
		"correo":"juanki@mail", 
		"telefono": "783278234", 
		"type_customer":1
		}
}
```
- **Response**
``` json
{
    "id": 1,
    "nombre": "jun carlos",
    "apellido": "jun carlos",
    "correo": "juanki@mail",
    "telefono": "783278234",
    "type_customer": 1
}
```

# Delete registry
- **Request**
- <code>localhost:8080/registry/1</code>