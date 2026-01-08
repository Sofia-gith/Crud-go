package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"encoding/json"
	


)

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetAge() int8
	GetPassword() string
	
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

type userDomain struct {
	Email    string  `json:"email"`
	Password string `json:"password"`	
	Name     string `json:"name"`
	Age      int8  `json:"age"`
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil{
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}


func(ud*userDomain) GetEmail()string{
	return ud.Email
}
func(ud*userDomain) GetName() string{
	return ud.Name
}
func(ud*userDomain) GetAge() int8{
	return ud.Age
}
func(ud*userDomain) GetPassword() string{
	return ud.Password
}

func (ud *userDomain) EncryptPassword(){
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password  = hex.EncodeToString(hash.Sum(nil))
}

