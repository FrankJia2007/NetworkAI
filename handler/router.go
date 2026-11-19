package handler

import (
    "net/http"

    "networkai/util"

    jwtMiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

var mySigningKey []byte

func InitRouter(config *util.TokenInfo) http.Handler {
    mySigningKey = []byte(config.Secret)

    jwtMiddleware := jwtMiddleware.New(jwtMiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })


    router := mux.NewRouter()

    router.Handle("/posts", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST")
    router.Handle("/posts/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET")
    // router.Handle("/posts/{id}", jwtMiddleware.Handler(http.HandlerFunc(deleteHandler))).Methods("DELETE")

    router.Handle("/users", http.HandlerFunc(signupHandler)).Methods("POST")
    router.Handle("/auth/login", http.HandlerFunc(signinHandler)).Methods("POST")

    originsOk := handlers.AllowedOrigins([]string{"*"})
    headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE"})

    return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}
