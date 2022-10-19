Тестовое задание на позицию стажёра-бэкендера

POST /moneyTransfer/clientBalance -- crediting funds to the balance

example:

{

    "userId" : 1,
    "creditAmount": 500
}

POST /fundReserve -- reserve funds

example:

{

    "userId" : 1,
    "serviceId": 2,
    "orderId": 2,
    "cost": 200,
}