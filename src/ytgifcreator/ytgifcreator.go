package ytgifcreator

import (
    "net/http"
    "net/rpc"
    "fmt"
    "html/template"
    "appengine"
    "reflect"
    "encoding/json"
)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/jsTemplates.html"))
var rpcConnection *rpc.Client

// A Context object for each request handler
type HandleContext struct {
    w http.ResponseWriter
    r *http.Request
    gaeContext appengine.Context
}

func (handleContext *HandleContext) RenderTemplate(template string, data interface{}) {
    err := templates.ExecuteTemplate(handleContext.w, template, data)
    if err != nil {
        http.Error(handleContext.w, err.Error(), http.StatusInternalServerError)
    }
}

// Application entry point
func init() {
    http.HandleFunc("/", entryHandler("IndexController"))
    http.HandleFunc("/_api/", entryHandler("ApiController"))
}

// Create a Controller based on its name
func getControllerByName(controllerName string, context HandleContext) Controller {
    switch controllerName {
    case "ApiController": return &ApiController { context }
    }

    return &IndexController { context }
}

// Create Request Handler with given Controller
func entryHandler(controllerName string) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        gaeContext := appengine.NewContext(r)
        handleContext := HandleContext{w, r, gaeContext}

        c := getControllerByName(controllerName, handleContext)
        c.Dispatch()
    }
}

// Try to call methodName on controller, or return 404 on fail
func FindMethodOr404 (ctrl Controller, methodName string) {
    requestMethod := reflect.ValueOf(ctrl).MethodByName(methodName);

    if requestMethod.IsValid() {
        requestMethod.Call(make([]reflect.Value, 0))
    } else {
        http.Error(ctrl.GetContext().w, "Api call not found", http.StatusNotFound)
    }
}

// Try to call rpc method
func RpcCall(serviceMethod string, args interface{}, reply interface{}) error {
    if rpcConnection == nil {
        var err error
        rpcConnection, err = rpc.DialHTTP("tcp", "localhost:8081")
        if err != nil {
            return fmt.Errorf("Can't find RPC Server.")
        }
    }

    done := <-rpcConnection.Go(serviceMethod, args, reply, nil).Done
    if done.Error != nil {
        return fmt.Errorf("Error in RPC Call: %#v, %v", rpcConnection.shutdown, done.Error.Error())
    }

    return nil
}

// Try to send given obj as json object to client
func JsonResponse(w http.ResponseWriter, obj interface{}) {
    data, err := json.Marshal(obj)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    w.Write(data)
}
