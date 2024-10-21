package model

type UserResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type UserCreateDto struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type UserUpdateDto struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	AvatarUrl string `json:"avatar_url"`
}

type UserDeleteDto struct {
	ID int64 `json:"id"`
}
