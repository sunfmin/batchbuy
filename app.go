package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
	"github.com/sunfmin/batchbuy/api"
	"github.com/sunfmin/batchbuy/model"
	"github.com/sunfmin/batchbuy/services"
	"github.com/sunfmin/batchbuy/services/apihttpimpl"
)

type Form map[string][]string

var serv = services.ServiceImpl{}
var appTemplate = template.New("appTemplate").Funcs(template.FuncMap{
	"newRow": func(index int) bool {
		return (index != 0 && index%3 == 0)
	},
	"formatTime": func(date string) string {
		if date == "" {
			return ""
		}
		dateT, err := services.StringToTime(date[:10])
		if err != nil {
			panic(err)
		}
		return dateT.Format(services.TimeFmt)
	},
})

var appRoot string

func init() {
	appRoot = os.Getenv("GOPATH") + "/src/github.com/sunfmin/batchbuy"

	appTemplate.ParseFiles([]string{
		appRoot + "/view/profile.html",
		appRoot + "/view/product.html",
		appRoot + "/view/order_list.html",
		appRoot + "/view/order.html",
		appRoot + "/view/user_list.html",
		appRoot + "/view/_app_header.html",
		appRoot + "/view/_app_footer.html",
	}...)
}

func main() {
	apihttpimpl.AddToMux("/api", http.DefaultServeMux)

	// handle assets and pages
	makeHandler("/", handleRootVisist)
	makeHandler("/favicon.ico", handleFavicon)
	makeHandler("/assets/", serveFile)
	makeHandler("/profile.html", profilePage)
	makeHandler("/product.html", productPage)
	makeHandler("/order.html", orderPage)
	makeHandler("/order_list.html", orderListPage)
	makeHandler("/user_list.html", userListPage)
	// makeHandler("/stop_order_today", )

	// handle api service
	handleProfile(serv)
	handleProduct(serv)
	handleOrder(serv)
	handleNoMoreOrderToday(serv)
	handleMakeMoreOrderToday(serv)
	handleIsNoMoreOrderToday(serv)

	println("Starting server on :8080")

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	defer model.End()
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	// TODO: find out whether there is a predefined variable like __FILE__ in ruby
	// path problem seems tricky in go, codes below only work when the package is invoked
	// on the GOPATH
	if strings.Contains(r.URL.Path, "/assets") {
		http.ServeFile(w, r, appRoot+"/"+r.URL.Path)
	} else {
		http.ServeFile(w, r, appRoot+"/view"+r.URL.Path)
	}
}

func profilePage(w http.ResponseWriter, r *http.Request) {
	err := appTemplate.ExecuteTemplate(w, "profile.html", "")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

// TODO: refactor three pages handler below: use multiple template files
func productPage(w http.ResponseWriter, r *http.Request) {
	products, err := serv.AllProducts()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	err = appTemplate.ExecuteTemplate(w, "product.html", products)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func orderPage(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Printf("Email Cookie: %s\n", err)

		http.Redirect(w, r, "/profile.html", http.StatusFound)
		return
	}

	user, err := model.GetUser(strings.Replace(email.Value, "%40", "@", 1))
	if err != nil {
		fmt.Printf("GetUser: %s\n", err)

		http.Redirect(w, r, "/profile.html", http.StatusTemporaryRedirect)
		return
	}

	dateM := r.Form["date"]
	var date time.Time
	if dateM == nil || dateM[0] == "" {
		date = time.Now()
	} else {
		date, err = services.StringToTime(dateM[0])
		if err != nil {
			fmt.Printf("services.StringToTime: %s\n", err)
			return
		}
	}
	today := date.Format(services.TimeFmt)

	pageVar := struct {
		// OrderedProducts []model.Product
		Orders                  []*api.Order
		MyYesterdayOrders       []*api.Order
		AvaliableProducts       []*api.Product
		MyTop3FavouriteProducts []*api.Product
		Top3PopularProducts     []*api.Product
		Weekday                 string
		Date                    string
		PreviousDay             string
		NextDay                 string
		IsNoMoreOrderToday      bool
	}{
		Date:        today,
		PreviousDay: date.AddDate(0, 0, -1).Format(services.TimeFmt),
		NextDay:     date.AddDate(0, 0, 1).Format(services.TimeFmt),
	}
	pageVar.AvaliableProducts, err = serv.MyAvaliableProducts(today, user.Email)
	pageVar.Top3PopularProducts, err = serv.Top3PopularProducts(today)
	pageVar.MyTop3FavouriteProducts, err = serv.MyTop3FavouriteProducts(user.Email, today)
	pageVar.Orders, err = serv.MyOrders(today, user.Email)
	pageVar.MyYesterdayOrders, err = serv.MyOrders(date.AddDate(0, 0, -1).Format(services.TimeFmt), user.Email)
	pageVar.IsNoMoreOrderToday, err = model.IsNoMoreOrderToday(time.Now())
	if err != nil {
		fmt.Printf("IsNoMoreOrderToday: %s\n", err)
		return
	}

	switch date.Weekday() {
	case 0:
		pageVar.Weekday = "星期天"
	case 1:
		pageVar.Weekday = "星期一"
	case 2:
		pageVar.Weekday = "星期二"
	case 3:
		pageVar.Weekday = "星期三"
	case 4:
		pageVar.Weekday = "星期四"
	case 5:
		pageVar.Weekday = "星期五"
	case 6:
		pageVar.Weekday = "星期六"
	}

	err = appTemplate.ExecuteTemplate(w, "order.html", pageVar)
	if err != nil {
		fmt.Printf("ExecuteTemplate: %s\n", err)
		return
	}
}

type orderHolder struct {
	Index   int
	Count   int
	Users   string
	Product string
}

func orderListPage(w http.ResponseWriter, r *http.Request) {
	dateM := r.Form["date"]
	var date time.Time
	var err error

	if dateM == nil || dateM[0] == "" {
		date = time.Now()
	} else {
		date, err = services.StringToTime(dateM[0])
		if err != nil {
			fmt.Printf("services.StringToTime: %s\n", err)
			return
		}
	}

	apiOrders, err := serv.OrderListOfDate(date.Format(services.TimeFmt))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	orders, ordersStr, brittanyOrdersStr := []orderHolder{}, "", ""
	for i, apiOrder := range apiOrders {
		var order orderHolder
		order.Index = i + 1
		order.Count = apiOrder.Count
		order.Product = apiOrder.Product.Name
		userNames := []string{}
		for _, user := range apiOrder.Users {
			count, err := model.GetOrderCount(user.Email, apiOrder.Product.Id, date)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			nameStr := fmt.Sprintf("%s (%d)", user.Name, count)
			userNames = append(userNames, nameStr)
		}
		order.Users = strings.Join(userNames, ", ")
		if strings.HasPrefix(order.Product, "Brittany") {
			brittanyOrdersStr += strings.Join([]string{order.Product, strconv.Itoa(order.Count)}, ", ")
			brittanyOrdersStr += ";\n"
		} else {
			ordersStr += strings.Join([]string{order.Product, strconv.Itoa(order.Count)}, ", ")
			ordersStr += ";\n"
		}

		orders = append(orders, order)
	}

	pageVar := struct {
		Orders             []orderHolder
		OrdersStr          string
		BrittanyOrdersStr  string
		Date               string
		PreviousDay        string
		NextDay            string
		UnorderedUsers     string
		IsNoMoreOrderToday bool
	}{
		Orders:            orders,
		OrdersStr:         ordersStr,
		BrittanyOrdersStr: brittanyOrdersStr,
		Date:              date.Format(services.TimeFmt),
		PreviousDay:       date.AddDate(0, 0, -1).Format(services.TimeFmt),
		NextDay:           date.AddDate(0, 0, 1).Format(services.TimeFmt),
	}

	unorderedUsers, err := model.UnorderedUsers(date)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// todayD := time.Now()
	// today := time.Date(todayD.Year(), todayD.Month(), todayD.Day(), 0, 0, 0, 0, todayD.Location())
	// fmt.Println(today)
	pageVar.IsNoMoreOrderToday, err = model.IsNoMoreOrderToday(time.Now())
	if err != nil {
		fmt.Printf("IsNoMoreOrderToday: %s\n", err)
		return
	}

	userLen := len(unorderedUsers)
	if userLen != 0 {
		for i, user := range unorderedUsers {
			pageVar.UnorderedUsers += user.Name
			if i != userLen-1 {
				pageVar.UnorderedUsers += ", "
			}
		}
	}

	err = appTemplate.ExecuteTemplate(w, "order_list.html", pageVar)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func userListPage(w http.ResponseWriter, r *http.Request) {
	users, err := serv.GetAllUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	err = appTemplate.ExecuteTemplate(w, "user_list.html", users)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

var decoder = schema.NewDecoder()

func handleProfile(service api.Service) {
	makeHandler("/profile", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			input := api.UserInput{}
			decoder.Decode(&input, r.Form)

			_, err := service.PutUser(input.Email, input)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			http.Redirect(w, r, "/order.html", http.StatusFound)
		case "DELETE":
			if err := service.RemoveUser(r.Form["email"][0]); err != nil {
				fmt.Printf("%s\n", err)
				return
			}
		}
	})
}

func handleNoMoreOrderToday(service api.Service) {
	makeHandler("/no_more_order_today", func(w http.ResponseWriter, r *http.Request) {
		date, err := services.StringToTime(r.Form["date"][0])
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		err = model.NoMoreOrderToday(date)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Fprintf(w, "Success")
	})
}

func handleMakeMoreOrderToday(service api.Service) {
	makeHandler("/make_more_order_today", func(w http.ResponseWriter, r *http.Request) {
		date, err := services.StringToTime(r.Form["date"][0])
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		err = model.MakeMoreOrderToday(date)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Fprintf(w, "Success")
	})
}

func handleIsNoMoreOrderToday(service api.Service) {
	makeHandler("/is_no_more_order_today", func(w http.ResponseWriter, r *http.Request) {
		date, err := services.StringToTime(r.Form["date"][0])
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		isNoMore, err := model.IsNoMoreOrderToday(date)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// data, err := json.Marshal(map[string]bool{"IsNoMoreOrderTdoay": isNoMore})
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Fprintf(w, fmt.Sprintf("%t", isNoMore))
	})
}

func handleProduct(service api.Service) {
	makeHandler("/product", func(w http.ResponseWriter, r *http.Request) {
		input := api.ProductInput{}
		err := decoder.Decode(&input, r.Form)
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		productId := string("")
		if len(r.Form["productid"]) != 0 {
			productId = r.Form["productid"][0]
		}
		product, err := service.PutProduct(productId, input)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		productBytes, err := json.Marshal(product)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Fprintf(w, string(productBytes))
	})
}

func handleOrder(service api.Service) {
	makeHandler("/order", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			count, err := strconv.Atoi(r.Form["count"][0])
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			_, err = service.PutOrder(r.Form["date"][0], r.Form["email"][0], r.Form["productid"][0], count)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			fmt.Fprintf(w, "Save Successfully")
		case "DELETE":
			err := service.RemoveOrder(r.Form["date"][0], r.Form["email"][0], r.Form["productid"][0])
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}

			fmt.Fprintf(w, "Delete Successfully")
		}
	})
}

func handleRootVisist(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/order.html", http.StatusFound)
}

func handleFavicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, appRoot+"/img/favicon.ico")
}

func makeHandler(path string, fn func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Printf("Form Parsing Error: %s\n", err)
			return
		}

		form := r.Form
		if path == "/assets/" {
			// fmt.Printf("Asset: %s\n\n", r.URL.Path)
		} else {
			fmt.Printf("Path: %s\nMethod: %s\nFormValus: %s\nCookies: %s\n\n", r.URL.Path, r.Method, form, r.Cookies())
		}

		fn(w, r)
	})
}
