# AddressBook

## Usecases

1. Как пользователь, я хочу иметь возможность просмотреть все мои контакты, чтобы использовать эту информацию.
1. Как пользователь, я хочу иметь возможность добавить контакт (Имя, фамилия, телефон, email, github), чтобы пополнять список моих контактов.
1. Как пользователь, я хочу иметь возможность изменить контакт, чтобы исправлять ошибки при добавлении.
1. Как пользователь, я хочу иметь возможность удалить контакт, чтобы не хранить неактульную информацию.

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
        "github" "user"
    }
```

Ответ: 202 Accepted
Location: /api/addressbook/contacts/1

### DELETE /api/addressbook/contacts/1

Ответ: 204 No Content
