package main

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}
