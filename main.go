package main // Обявляем пакет main
import (
	"fmt"
	SpaceProxy "proxy/library"
)

func main() {
	apikey := "FKUNM9gqxcCIzdC1KWdWjDyQishplhLDivA9N0Wb"
	mp := make(map[string]interface{})

	/* Метод для получения списка прокси. */
	//countries := []string{}
	//orders := []string{}
	//mp = SpaceProxy.Proxies(apikey,
	//	"",
	//	"",
	//	"",
	//	countries,
	//	"",
	//	"",
	//	"",
	//	"",
	//	orders,
	//	"",
	//	"")
	//fmt.Println(mp, "\n")

	/* Метод для получения списка заказов. */
	mp = SpaceProxy.Orders(apikey, "", "")
	fmt.Println(mp, "\n")

	/* Продление прокси. Оплата со счета. Метод POST. */
	//mp = SpaceProxy.Renew(apikey, "1586231", "5", "")
	//fmt.Println(mp, "\n")

	/* Создание нового заказа. Оплата со счета. Метод POST. */
	//mp = SpaceProxy.NewOrder(apikey, "dedicated", "4", "ru", "1", "5", "", "48743")
	//fmt.Println(mp, "\n")

	/* Получить список стран с вложенными городами. Код страны и ID города можно использовать для фильтрации ips. */
	//mp = SpaceProxy.Countries(apikey)
	//fmt.Println(mp, "\n")

	/* Получить список IP по параметрам. ID используются в методах создания заказа и продления. */
	//mp = SpaceProxy.Ips(apikey, "dedicated", "4", "ru")
	//fmt.Println(mp, "\n")

	/* Получить количество IP по параметрам. Если количество больше 1000, будет показано как 1000. */
	//mp = SpaceProxy.IpsCount(apikey, "dedicated", "4", "ru")
	//fmt.Println(mp, "\n")

	/* Получение суммы заказа до оплаты. Метод POST. */
	//ips := []string{"48743", "48747"}
	//mp = SpaceProxy.NewOrderAmount(apikey, "dedicated", "4", "ru", "1", "5", "", ips)
	//fmt.Println(mp, "\n")

	/* Получение остатка на балансе. Обычном и партнерском. */
	//mp = SpaceProxy.Balance(apikey)
	//fmt.Println(mp, "\n")
}
