package main

import (
    "github.com/sunfmin/batchbuy/api"
    "fmt"
    "github.com/gorilla/schema"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "strings"
    "time"
)

type Form map[string][]string

var controller = Controller{}

func main() {
    // handle assets and pages
    makeHandler("/assets/", serverFile)
    makeHandler("/profile.html", serverFile)
    makeHandler("/product.html", productPage)
    makeHandler("/order.html", orderPage)
    makeHandler("/order_list.html", orderListPage)

    // handle api service
    handleProfile(controller)
    handleProduct(controller)
    handleOrder(controller)
    // handleProductAll(controller)

    s := &http.Server{
        Addr: ":8080",
        // Handler:        myHandler,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
}

const appRoot = "/usr/local/go/src/pkg/github.com/sunfmin/batchbuy"

func serverFile(w http.ResponseWriter, r *http.Request) {
    // todo find out whether there is a predefined variable like __FILE__ in ruby
    // path problem seems tricky in go, codes below only work when the package is invoked
    // on the root path of this workspace
    if strings.Contains(r.URL.Path, "/assets") {
        http.ServeFile(w, r, appRoot+"/"+r.URL.Path)
    } else {
        http.ServeFile(w, r, appRoot+"/view"+r.URL.Path)
    }
}

// TODO: refactor three pages handler below: use multiple template files
func productPage(w http.ResponseWriter, r *http.Request) {
    products, _ := controller.AllProducts()
    var templates = template.Must(template.ParseFiles(appRoot + "/view/product.html"))

    templates.ExecuteTemplate(w, "product.html", products)
}

func orderPage(w http.ResponseWriter, r *http.Request) {
    products, _ := controller.ProductListOfDate(time.Now().Format(timeFmt))
    var templates = template.Must(template.ParseFiles(appRoot + "/view/order.html"))

    templates.ExecuteTemplate(w, "order.html", products)
}

type orderHolder struct {
    Index   int
    Count   int
    Users   string
    Product string
}

func orderListPage(w http.ResponseWriter, r *http.Request) {
    apiOrders, _ := controller.OrderListOfDate(time.Now().Format(timeFmt))

    orders := []orderHolder{}
    var order orderHolder
    for i, apiOrder := range apiOrders {
        order = orderHolder{}
        order.Index = i + 1
        order.Count = apiOrder.Count
        order.Product = apiOrder.Product.Name
        userNames := []string{}
        for _, user := range apiOrder.Users {
            userNames = append(userNames, user.Name)
        }
        order.Users = strings.Join(userNames, ", ")

        orders = append(orders, order)
    }
    fmt.Printf("%s", orders)

    var templates = template.Must(template.ParseFiles(appRoot + "/view/order_list.html"))
    templates.ExecuteTemplate(w, "order_list.html", orders)
}

var decoder = schema.NewDecoder()

func handleProfile(service api.Service) {
    makeHandler("/profile", func(w http.ResponseWriter, r *http.Request) {
        input := api.UserInput{}
        decoder.Decode(&input, r.Form)

        service.PutUser(input.Email, input)

        http.Redirect(w, r, "/order.html", http.StatusFound)
    })
}

func handleProduct(service api.Service) {
    makeHandler("/product", func(w http.ResponseWriter, r *http.Request) {
        input := api.ProductInput{}
        decoder.Decode(&input, r.Form)
        fmt.Printf("%s\n%s\n", r.Form["id"][0], input)
        service.PutProduct(r.Form["id"][0], input)

        http.Redirect(w, r, "/product.html", http.StatusFound)
    })
}

func handleOrder(service api.Service) {
    makeHandler("/order", func(w http.ResponseWriter, r *http.Request) {
        count, _ := strconv.Atoi(r.Form["count"][0])
        service.PutOrder(time.Now().Format(timeFmt), r.Form["email"][0], r.Form["productid"][0], count)

        http.Redirect(w, r, "/order_list.html", http.StatusFound)
    })
}

// func handleProductAll(service api.Service) {
// 	makeHandler("/product/all", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("%s", r.URL.Path)
// 		
// 		http.Redirect(w, r, "/product.html", http.StatusFound)
// 	})
// }

func makeHandler(path string, fn func(http.ResponseWriter, *http.Request)) {
    http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        form := r.Form
        fmt.Printf("Path: %s\nFormValus: %s\n\n", r.URL.Path, form)

        fn(w, r)
    })
}
