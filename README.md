

main.go запуск и меню

    users.go
        NewAccaunt: создание аккаунта
        AllUsers: список пользователей
cards
    NewCards.go
        NewCard: создание дебетовой карты
        NewCreditCard: создание кредитки
        CreditLimit: расчёт лимита кредитки

    Tranzaction.go
        Card.UpDep, CreditCard.UpDep: пополнение
        UpDepCard: выбор карты для пополнении
        Card.DownDep, CreditCard.DownDep: снятие
        DownDepCard: выбор карты для снятия

    ViewCard.go
        Card.ViewCard, CreditCard.ViewCard: информация о карте
        PrintCards: вывод карты 
        ViewUserCards: карты пользователя.

users
    TransID.go
        TranzactionForId: перевод по ID пользователей
    
    Users.go
        NewAccaunt: новый пользователь
        Allusers: все пользователи 

structs
    Structs.go структуры данных
