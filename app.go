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
        dateT, err := stringToTime(date[:10])
        if err != nil {
            panic(err)
        }
        return dateT.Format(timeFmt);
    },
})    

const appRoot = "src/github.com/sunfmin/batchbuy"

func init() {
    appTemplate.ParseFiles([]string{appRoot + "/view/profile.html", appRoot + "/view/product.html", appRoot + "/view/order_list.html", appRoot + "/view/order.html", appRoot + "/view/_app_header.html", appRoot + "/view/_app_footer.html"}...)
}

func main() {
    // handle assets and pages
    makeHandler("/", handleRootVisist)
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
    products, err := controller.AllProducts()
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
        return
    }
    
    dateM := r.Form["date"]
    var date time.Time
    if dateM == nil || dateM[0] == "" {
        date = time.Now()
    } else {
        date, err = stringToTime(dateM[0])
        if err != nil {
            fmt.Printf("stringToTime: %s\n", err)
            return
        }
    }
    
    pageVar := struct {
        // OrderedProducts []model.Product
        Orders []*api.Order
        AvaliableProducts []model.Product
        Date string
        PreviousDay string
        NextDay string
    }{
        Date: date.Format(timeFmt), 
        PreviousDay: date.AddDate(0, 0, -1).Format(timeFmt),
        NextDay: date.AddDate(0, 0, 1).Format(timeFmt),
    }
    pageVar.AvaliableProducts, err = user.AvaliableProducts(date)
    pageVar.Orders, err = user.OrdersForApi(date)
    
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
        date, err = stringToTime(dateM[0])
        if err != nil {
            fmt.Printf("stringToTime: %s\n", err)
            return
        }
    }
    
    apiOrders, err := controller.OrderListOfDate(date.Format(timeFmt))
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

    orders, ordersStr := []orderHolder{}, ""
    var order orderHolder
    for i, apiOrder := range apiOrders {
        order = orderHolder{}
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
        ordersStr += strings.Join([]string{strconv.Itoa(order.Index), order.Product, strconv.Itoa(order.Count)}, ", ")
        ordersStr += ";\n"
        
        orders = append(orders, order)
    }
    
    pageVar := struct {
        Orders []orderHolder
        OrdersStr string
        Date string
        PreviousDay string
        NextDay string
    }{
        Orders: orders,
        OrdersStr: ordersStr,
        Date: date.Format(timeFmt), 
        PreviousDay: date.AddDate(0, 0, -1).Format(timeFmt),
        NextDay: date.AddDate(0, 0, 1).Format(timeFmt),
    }
    
    err = appTemplate.ExecuteTemplate(w, "order_list.html", pageVar)
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
}

var decoder = schema.NewDecoder()

func handleProfile(service api.Service) {
    makeHandler("/profile", func(w http.ResponseWriter, r *http.Request) {
        input := api.UserInput{}
        decoder.Decode(&input, r.Form)

        _, err := service.PutUser(input.Email, input)
        if err != nil {
            fmt.Printf("%s\n", err)
            return
        }

        http.Redirect(w, r, "/order.html", http.StatusFound)
    })
}

func handleProduct(service api.Service) {
    makeHandler("/product", func(w http.ResponseWriter, r *http.Request) {
        input := api.ProductInput{}
        err := decoder.Decode(&input, r.Form)
        if err != nil { fmt.Printf("%s\n", err) }
        
        productId := string("")
        if (len(r.Form["productid"]) != 0) {
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

func makeHandler(path string, fn func(http.ResponseWriter, *http.Request)) {
    http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
        err := r.ParseForm()
        if err != nil {
            fmt.Printf("%s\n", err)
            return
        }
        
        form := r.Form
        if path == "/assets/" {
            fmt.Printf("Asset: %s\n\n", r.URL.Path)
        } else {
            fmt.Printf("Path: %s\nMethod: %s\nFormValus: %s\nCookies: %s\n\n", r.URL.Path, r.Method, form, r.Cookies())            
        }

        fn(w, r)
    })
}