# AddressBook

## Usecases

1. Как пользователь, я хочу иметь возможность просмотреть все мои контакты, чтобы использовать эту информацию.
1. Как пользователь, я хочу иметь возможность добавить контакт (имя, фамилия, телефон, email, github), чтобы пополнять список моих контактов.
1. Как пользователь, я хочу иметь возможность изменить контакт, чтобы исправлять ошибки при добавлении.
1. Как пользователь, я хочу иметь возможность удалить контакт, чтобы не хранить неактуальную информацию.

## REST API

### GET /api/addressbook/contacts

Ответ: 200 OK
```json
    [{
        "name": "Имя",
        "surname": "Фамилия",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "github" "user"
    }]
```

### POST /api/addressbook/contacts

Тело запроса:

```json
    {
        "name": "Имя",
        "surname": "Фамилия",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "github" "user"
    }
```

Ответ: 201 Created
Location: /api/addressbook/contacts/1

### PUT /api/addressbook/contacts/1

Тело запроса:

```json
    {
        "name": "Имя",
        "surname": "Фамилия",
        "phone": "+7-999-999-99-99",
        "email": "user@domain.ru",
        "github": "user"
    }
```

Ответ: 202 Accepted
Location: /api/addressbook/contacts/1

### DELETE /api/addressbook/contacts/1

Ответ: 204 No Content

## Как собрать и запустить

Backend:

```bat
cd backend
docker build -f Dockerfile -t addressbookbackend:<имя ветки> .
docker run --rm --name addressbookbackend -e NAME=<параметр приложения> addressbookbackend:<имя ветки>
```

Frontend:

```bat
cd frontend
docker build -f Dockerfile -t addressbookfrontend:<имя ветки> .
docker run -d --rm --name addressbookfrontend -p 80:80 addressbookfrontend:<имя ветки>
```