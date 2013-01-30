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
    "github.com/sunfmin/batchbuy/model"
    "encoding/json"
)

type Form map[string][]string

var controller = Controller{}
var appTemplate = template.New("appTemplate").Funcs(template.FuncMap{
    "newRow": func(index int) bool { 
        return (index != 0 && index % 3 == 0);
    },
    "formatTime": func(date string) string {
        if date == "" { 
            return "" 
        }
        return stringToTime(date[:10]).Format(timeFmt);
    },
})    

const appRoot = "src/github.com/sunfmin/batchbuy"

func init() {
    appTemplate.ParseFiles([]string{appRoot + "/view/profile.html", appRoot + "/view/product.html", appRoot + "/view/order_list.html", appRoot + "/view/order.html", appRoot + "/view/_app_header.html", appRoot + "/view/_app_footer.html"}...)
}

func main() {
    // handle assets and pages
    makeHandler("/assets/", serveFile)
    makeHandler("/profile.html", profilePage)
    makeHandler("/product.html", productPage)
    makeHandler("/order.html", orderPage)
    makeHandler("/order_list.html", orderListPage)

    // handle api service
    handleProfile(controller)
    handleProduct(controller)
    handleOrder(controller)
    
    s := &http.Server{
        Addr: ":8080",
        // Handler:        myHandler,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
    
    // model.ConnectDb()
    defer model.End()
}

func serveFile(w http.ResponseWriter, r *http.Request) {
    // TODO: find out whether there is a predefined variable like __FILE__ in ruby
    // path problem seems tricky in go, codes below only work when the package is invoked
    // on the root path of this workspace
    if strings.Contains(r.URL.Path, "/assets") {
        http.ServeFile(w, r, appRoot+"/"+r.URL.Path)
    } else {
        http.ServeFile(w, r, appRoot+"/view"+r.URL.Path)
    }
}

func profilePage(w http.ResponseWriter, r *http.Request) {
    appTemplate.ExecuteTemplate(w, "profile.html", "")
}

// TODO: refactor three pages handler below: use multiple template files
func productPage(w http.ResponseWriter, r *http.Request) {
    products, _ := controller.AllProducts()

    appTemplate.ExecuteTemplate(w, "product.html", products)
}

func orderPage(w http.ResponseWriter, r *http.Request) {
    products, _ := controller.ProductListOfDate(time.Now().Format(timeFmt))
    
    appTemplate.ExecuteTemplate(w, "order.html", products)
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

    appTemplate.ExecuteTemplate(w, "order_list.html", orders)
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
        // fmt.Printf("%s\n%s\n", r.Form["productid"][0], input)
        
        productId := string("")
        if (len(r.Form["productid"]) != 0) {
            productId = r.Form["productid"][0]
        }
        product, _ := service.PutProduct(productId, input)

        // http.Redirect(w, r, "/product.html", http.StatusFound)
        productBytes, _ := json.Marshal(product)
        fmt.Fprintf(w, string(productBytes))
    })
}

func handleOrder(service api.Service) {
    makeHandler("/order", func(w http.ResponseWriter, r *http.Request) {
        count, _ := strconv.Atoi(r.Form["count"][0])
        service.PutOrder(r.Form["date"][0], r.Form["email"][0], r.Form["productid"][0], count)

        http.Redirect(w, r, "/order_list.html", http.StatusFound)
    })
}

func makeHandler(path string, fn func(http.ResponseWriter, *http.Request)) {
    http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        form := r.Form
        fmt.Printf("Path: %s\nMethod: %s\nFormValus: %s\n\n", r.URL.Path, r.Method, form)

        fn(w, r)
    })
}