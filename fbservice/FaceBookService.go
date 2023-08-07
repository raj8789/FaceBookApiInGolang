package fbservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	fb "github.com/huandu/facebook/v2"
)

const (
	access_token string = "EAAJZA2ZC0aPhQBAIilFr5giD1ZCpfgn7HsnfcRcLp9wGhjhhhiugybA34ZBODaZBZB2y7E0Uj1gk8dnKBXnrdD6opTSAFKiynZC8DNPFdDbNvlx1DtsJugsXQ4eZBZCYxaGz3L77ZCx94NhZCgdYlevPFBoZCDBl4gIvcWqowVbRgZCx5yt3qbmJTV8Hve1KpDfkOZBvYiAlv8hD8lrYoT9f4PWZBRnGZAu8i50615y6q7ZCV3aDfqj7bdHUv3r9xHoCh6rX5q"
)

func GetUserName(rw http.ResponseWriter, re *http.Request) {
	res, _ := fb.Get("/me", fb.Params{
		"fields":       "name",
		"access_token": access_token,
	})
	fmt.Printf("Name :=%v", res["name"])
	m := make(map[string]interface{})
	m["Name"] = res["name"]
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
	fmt.Println("")
}
func GetUserEmail(rw http.ResponseWriter, re *http.Request) {
	res, _ := fb.Get("/me", fb.Params{
		"fields":       "email",
		"access_token": access_token,
	})
	fmt.Printf("Email :=%v", res["email"])
	m := make(map[string]interface{})
	m["email"] = res["email"]
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
	fmt.Println("")
}
func GetUserDeatils(rw http.ResponseWriter, re *http.Request) {
	res, _ := fb.Get("/me", fb.Params{
		"fields":       "email,name,birthday,gender,hometown,friends",
		"access_token": access_token,
	})
	strArr := [6]string{"friends", "hometown", "gender", "birthday", "name", "email"}
	m := make(map[string]interface{})
	for _, value := range strArr {
		m[value] = res[value]
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
	fmt.Println("")
}
func PostPicture(rw http.ResponseWriter, re *http.Request) {
	//imagePath := `C:\Users\raush\OneDrive\Pictures\Screenshots\test.png`
    caption := "MSD Post Match Presentation Pic posted on My Timeline  Through My FaceBook Api In Golang"
	pageID := "103531532820275"
    // Upload the photo
	postParams := fb.Params{
        "message": caption,
		"access_token": access_token,
		//"source": fb.File(imagePath),
		//"url":imagePath,
    }

    res, err := fb.Post("/"+pageID+"/feed", postParams)

    if err != nil {
        fmt.Printf("Error posting picture: %v\n", err)
        return
    }

    fmt.Println("Picture posted successfully!")
    fmt.Println("Photo ID:", res["id"])
}

