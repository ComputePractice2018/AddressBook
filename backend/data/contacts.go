package data

// Contact структура для хранения записи адресной книги
type Contact struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Github  string `json:"github"`
}

// ContactList хрангимый список контактов
var ContactList = []Contact{Contact{
	Name:    "Имя",
	Surname: "Фамилия",
	Phone:   "+7-999-999-99-99",
	Email:   "mail@domain.ru",
	Github:  "user"}}
