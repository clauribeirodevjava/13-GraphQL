# GraphQL <img src="https://ih1.redbubble.net/image.520470372.9907/raf,750x1000,075,t,oatmeal_heather.u4.jpg" width="100" height="100">

## Gerando o Esqueleto do servidor graphiQL

**Entrar no site**


[https://gqlgen.com/](https://gqlgen.com/)

**Add github.com/99designs/gqlgen to your project’s tools.go**
```
printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go 
```
```
go mod tidy 	
```
**Initialise gqlgen config and generate models**
```go run github.com/99designs/gqlgen init```

**Start the graphql server**
```
go run server.go
```

**Gerar schemas Graphql**


# Criar os esquemas:

```go run github.com/99designs/gqlgen generate```

## Separar os models para poder gerar os relacionamentos :


**category.go**
```go
package model

type Category struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description *string `json:"description,omitempty"`
}
```


course.go
```go
package model

type Course struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description *string `json:"description,omitempty"`
}
```

**Implementar os BDS :**

**Implementar Resolvers:**
