package main

import "Crud-go/src/configuration/rest_err"

func test(message string) *rest_err.RestErr {
    return rest_err.NewBadRequestError(message)
}