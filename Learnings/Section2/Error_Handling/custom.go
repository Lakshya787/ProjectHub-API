package main

import "fmt"

//type error interface{
//	Error() string
//}

type CustomError struct{
	Code int
	Message string
}

func (e CustomError) Error() string{
	return  fmt.Sprintf("Error code %d with message: %s",e.Code,e.Message)
}

func check_(code int) (string,error){
	if code ==400{
		return "",&CustomError{Code:400,Message: "Invalid request"}
	}
	return "success",nil
}

func custom(){
	if _,err:=check(400);err!=nil{
		fmt.Println(err)
	}
}