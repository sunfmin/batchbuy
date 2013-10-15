// Generated by github.com/hypermusk/hypermusk
// DO NOT EDIT

package apihttpimpl

import (
	"time"
	"encoding/json"
	"github.com/sunfmin/batchbuy/api"
	"github.com/sunfmin/batchbuy/services"
	"net/http"
)

var _ = time.Sunday

type CodeError interface {
	Code() string
}

type SerializableError struct {
	Code    string
	Message string
	Reason  error
}

func (s *SerializableError) Error() string {
	return s.Message
}

func NewError(err error) (r error) {
	se := &SerializableError{Message:err.Error()}
	ce, yes := err.(CodeError)
	if yes {
		se.Code = ce.Code()
	}
	se.Reason = err
	r = se
	return
}

func AddToMux(prefix string, mux *http.ServeMux) {
	
	mux.HandleFunc(prefix+"/Service/PutProduct.json", Service_PutProduct)
	mux.HandleFunc(prefix+"/Service/RemoveProduct.json", Service_RemoveProduct)
	mux.HandleFunc(prefix+"/Service/PutUser.json", Service_PutUser)
	mux.HandleFunc(prefix+"/Service/RemoveUser.json", Service_RemoveUser)
	mux.HandleFunc(prefix+"/Service/PutOrder.json", Service_PutOrder)
	mux.HandleFunc(prefix+"/Service/RemoveOrder.json", Service_RemoveOrder)
	mux.HandleFunc(prefix+"/Service/ProductListOfDate.json", Service_ProductListOfDate)
	mux.HandleFunc(prefix+"/Service/OrderListOfDate.json", Service_OrderListOfDate)
	mux.HandleFunc(prefix+"/Service/MyAvaliableProducts.json", Service_MyAvaliableProducts)
	mux.HandleFunc(prefix+"/Service/MyOrders.json", Service_MyOrders)
	return
}


var service api.Service = services.DefaultService

type ServiceData struct {
}


type Service_PutProduct_Params struct {
	Params struct {
		Id string
		Input api.ProductInput
	}
}

type Service_PutProduct_Results struct {
	Product *api.Product
	Err error

}

func Service_PutProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_PutProduct_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_PutProduct_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Product, result.Err = s.PutProduct(p.Params.Id, p.Params.Input)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_RemoveProduct_Params struct {
	Params struct {
		Id string
	}
}

type Service_RemoveProduct_Results struct {
	Err error

}

func Service_RemoveProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_RemoveProduct_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_RemoveProduct_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Err = s.RemoveProduct(p.Params.Id)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_PutUser_Params struct {
	Params struct {
		Email string
		Input api.UserInput
	}
}

type Service_PutUser_Results struct {
	User *api.User
	Err error

}

func Service_PutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_PutUser_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_PutUser_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.User, result.Err = s.PutUser(p.Params.Email, p.Params.Input)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_RemoveUser_Params struct {
	Params struct {
		Email string
	}
}

type Service_RemoveUser_Results struct {
	Err error

}

func Service_RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_RemoveUser_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_RemoveUser_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Err = s.RemoveUser(p.Params.Email)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_PutOrder_Params struct {
	Params struct {
		Date string
		Email string
		ProductId string
		Count int
	}
}

type Service_PutOrder_Results struct {
	Order *api.Order
	Err error

}

func Service_PutOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_PutOrder_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_PutOrder_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Order, result.Err = s.PutOrder(p.Params.Date, p.Params.Email, p.Params.ProductId, p.Params.Count)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_RemoveOrder_Params struct {
	Params struct {
		Date string
		Email string
		ProductId string
	}
}

type Service_RemoveOrder_Results struct {
	Err error

}

func Service_RemoveOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_RemoveOrder_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_RemoveOrder_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Err = s.RemoveOrder(p.Params.Date, p.Params.Email, p.Params.ProductId)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_ProductListOfDate_Params struct {
	Params struct {
		Date string
	}
}

type Service_ProductListOfDate_Results struct {
	Products []*api.Product
	Err error

}

func Service_ProductListOfDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_ProductListOfDate_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_ProductListOfDate_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Products, result.Err = s.ProductListOfDate(p.Params.Date)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_OrderListOfDate_Params struct {
	Params struct {
		Date string
	}
}

type Service_OrderListOfDate_Results struct {
	Orders []*api.Order
	Err error

}

func Service_OrderListOfDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_OrderListOfDate_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_OrderListOfDate_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Orders, result.Err = s.OrderListOfDate(p.Params.Date)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_MyAvaliableProducts_Params struct {
	Params struct {
		Date string
		Email string
	}
}

type Service_MyAvaliableProducts_Results struct {
	Products []*api.Product
	Err error

}

func Service_MyAvaliableProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_MyAvaliableProducts_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_MyAvaliableProducts_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Products, result.Err = s.MyAvaliableProducts(p.Params.Date, p.Params.Email)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}

type Service_MyOrders_Params struct {
	Params struct {
		Date string
		Email string
	}
}

type Service_MyOrders_Results struct {
	Orders []*api.Order
	Err error

}

func Service_MyOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	var p Service_MyOrders_Params
	if r.Body == nil {
		panic("no body")
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&p)
	var result Service_MyOrders_Results
	enc := json.NewEncoder(w)
	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}

	s := service

	if err != nil {
		result.Err = NewError(err)
		enc.Encode(result)
		return
	}
	result.Orders, result.Err = s.MyOrders(p.Params.Date, p.Params.Email)
	if result.Err != nil {
		result.Err = NewError(result.Err)
	}
	err = enc.Encode(result)
	if err != nil {
		panic(err)
	}
	return
}








