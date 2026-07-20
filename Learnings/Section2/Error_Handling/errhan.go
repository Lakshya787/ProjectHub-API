package main

import (
	"errors"
	"fmt"
)

var ErrInvalidFile=errors.New("File is not valid")

func check(num int) (string, error) {
	if num < 0 {
		return "", errors.New("Number is negative")
	}
	return "Number is positive",nil
}

func read(filename string) (error){
	if filename==""{
		return ErrInvalidFile
	}
	return nil
}

func file () (string,error){
	e:=read("")
	if e!=nil{
		return "",fmt.Errorf("Error reading file: %w",e)
	}
	return "File exists",nil
}

func errhan(){
	_,e:=check(-1)
	if e!=nil{
		fmt.Println(e)
	}

	_,err:=file()
	if err!=nil{
		fmt.Println(err)
		we:=errors.Unwrap(err)
		if errors.Is(we,ErrInvalidFile){
			fmt.Println(we)
		}
	}
}