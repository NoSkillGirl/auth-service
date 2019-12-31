package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/NoSkillGirl/auth-service/models"
	userController "github.com/NoSkillGirl/user-service/controllers"
	userModel "github.com/NoSkillGirl/user-service/models"
)

//UserLoginRequest struct
type UserLoginRequest struct {
	Name     string
	Password string
}

//UserRegisterRequest struct
type UserRegisterRequest struct {
	ID       int
	Name     string
	PhoneNo  string `json:"phone_no"`
	EmailID  string `json:"email_id"`
	Password string
}

//UserLogin function
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var reqJSON UserLoginRequest
	fmt.Println(r)
	err := json.NewDecoder(r.Body).Decode(&reqJSON)

	if err != nil {
		panic(err)
	}
	log.Println(reqJSON)

	u, _ := userModel.UserExist(reqJSON.Name, reqJSON.Password)

	data := userController.ResponseV2{
		Status: 200,
		Response: userController.ResponseMsgV2{
			User: u,
		},
		Error: userController.ErrorMessage{},
	}

	if len(u) > 0 {
		authCode := models.AddAuthCode(u[0].ID)

		data.Response.Msg = authCode
	} else {
		data.Response.Msg = "User Not Found"
	}

	json.NewEncoder(w).Encode(data)
	// json.NewEncoder(w).Encode(authCode)
	// old - yours
	// urlStr := "http://localhost:8082/searchUser"
	// // Build out the data for our message
	// v := url.Values{}
	// v.Set("name", reqJSON.Name)
	// v.Set("password", reqJSON.Password)
	// rb := *strings.NewReader(v.Encode())
	// // Create client
	// client := &http.Client{}
	// req, _ := http.NewRequest("POST", urlStr, &rb)
	// resp, _ := client.Do(req)
	// fmt.Println(resp)
	// old - mine
	// // url := "http://restapi3.apiary.io/notes"
	// url := "http://localhost:8082/searchUser"
	// fmt.Println("URL:>", url)
	// b, err := json.Marshal(reqJSON)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println()
	// var jsonStr = []byte(string(b))
	// // var jsonStr = r.Body
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// // req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

}

//RegisterUser function
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var reqJSON UserRegisterRequest

	err := json.NewDecoder(r.Body).Decode(&reqJSON)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	fmt.Println(reqJSON)
	//log.Println(reqJSON)
	//errOccured, duplicateExist := userModel.AddUser(reqJSON.Name, reqJSON.PhoneNo, reqJSON.EmailID, reqJSON.Password)
	errOccured, duplicateExist := userModel.AddUser(reqJSON.Name, reqJSON.PhoneNo, reqJSON.EmailID, reqJSON.Password)
	if errOccured == true {
		data := userController.Response{
			Status:   500,
			Response: userController.ResponseMsg{},
			Error: userController.ErrorMessage{
				Msg: "Internal server Error",
			},
		}
		json.NewEncoder(w).Encode(data)
	} else if errOccured == false && duplicateExist == true {
		data := userController.Response{
			Status:   500,
			Response: userController.ResponseMsg{},
			Error: userController.ErrorMessage{
				Msg: "User already exist",
			},
		}
		json.NewEncoder(w).Encode(data)
	} else {
		data := userController.Response{
			Status: 200,
			Response: userController.ResponseMsg{
				Msg: "User succesfully created",
			},
			Error: userController.ErrorMessage{},
		}
		json.NewEncoder(w).Encode(data)
	}
	//json.NewEncoder(w).Encode(data)
}
