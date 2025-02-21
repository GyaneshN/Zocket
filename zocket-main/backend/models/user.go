package models


type User struct {
    ID       string `json:"id"`
    Email    string `json:"email"`
     Password string `json:"password"`
    Name     string `json:"name"`
    TeamID   string `json:"team_id"`
}