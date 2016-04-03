package handle

import (
	"net/http"
	"encoding/json"
	"strconv"
	"log"

	v "../variables"
	stc "../structures"
	connect "../connect"

	"github.com/go-martini/martini"
)

func AddNewUser(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
    var user stc.UserRequest   
    err := decoder.Decode(&user)
   	if err != nil {
       	panic("Is not possible read the JSON request")
   	}

	errors := stc.Errors{}
	status := v.Success
	if connect.AddUser(user.User_Name) != true{
		status = v.Error
       	errors = append(errors, stc.Error{ Body: "Is not possible add the user"})
	}
	response := stc.Result { Status: status, Data: stc.Users{} ,Errors : errors}
    json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, params martini.Params){
	user_id := params["user_id"]
	user_name := params["user_name"]

	errors := stc.Errors{}
	status, error := IsNumber(user_id)
	if status == v.Success{
		if connect.UpdateUser(user_id, user_name) != true{
			status = v.Error
	     	errors = append(errors, stc.Error{ Body: "Is not possible update the user"})
		}
	}else{
		errors = append(errors, error)
	}
	response := stc.Result { Status: status, Data: stc.Users{} ,Errors : errors}
	json.NewEncoder(w).Encode(response)
}

func UpdateUserJson(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	log.Println(r.Body)
    var user stc.UserRequest   
    err := decoder.Decode(&user)
   	if err != nil {
       	panic("Is not possible read the JSON request")
   	}
	errors := stc.Errors{}
	log.Println(user.User_Id)
	status, error := IsNumber(user.User_Id)
	if status == v.Success{
		if connect.UpdateUser(user.User_Id, user.User_Name) != true{
			status = v.Error
	     	errors = append(errors, stc.Error{ Body: "Is not possible update the user"})
		}
	}else{
		errors = append(errors, error)
	}
	response := stc.Result { Status: status, Data: stc.Users{} ,Errors : errors}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, params martini.Params){	
	user_id := params["user_id"]
	errors := stc.Errors{}
	status, error := IsNumber(user_id)
	if status == v.Success{
		if connect.DeleteUser(user_id) != true{
			status = v.Error
       		errors = append(errors, stc.Error{ Body: "Is not possible delete the user"})
		}
	}else{
		errors = append(errors, error)
	}
	response := stc.Result { Status: status, Data: stc.Users{} ,Errors : errors}
    json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request, params martini.Params){	
	user_id := params["user_id"]
	errors := stc.Errors{}
	data := stc.Users{}
	status, error := IsNumber(user_id)
	if status == v.Success{
		data = connect.GetUser(user_id)
		if (len(data) == 0){
      		status = v.Error
       		errors = append(errors, stc.Error{ Body: "User not found"})
    	}
	}else{
		errors = append(errors, error)
	}
	response := stc.Result { Status : status, Data: data ,Errors : errors}
   	json.NewEncoder(w).Encode(response)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	errors := stc.Errors{}			
	data := connect.GetAllUser()

	response := stc.Result { Status : v.Success, Data: data, Errors : errors}
    json.NewEncoder(w).Encode(response)
}

func IsNumber(number string) (status int, error stc.Error) {
	_, err := strconv.Atoi(number)
    if err != nil {
        return v.Error, stc.Error{ Body: "The param is not a number"}
    }
	return v.Success, stc.Error{}
}
