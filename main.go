package main

import (
        "fmt"
        "log"
        "net/http"
        "io/ioutil"
        "encoding/json"

        "github.com/gorilla/mux"
)

const (
    MethodReceiver = "VolumeDriver"
)

func main() {
        address := []byte("http://127.0.0.1:9080")
        err := ioutil.WriteFile(
            "/usr/share/docker/plugins/isilon.spec", address, 0777)
        if err != nil {
            log.Fatal(err)
        }

        router := mux.NewRouter()
        router.NotFoundHandler = http.HandlerFunc(notFound)

        router.Methods("POST").Path("/Plugin.Activate").HandlerFunc(handshake)
        log.Println("Created handshake handler")

        handleMethod := func(method string, h http.HandlerFunc) {
                router.Methods("POST").Path(fmt.Sprintf("/%s.%s", MethodReceiver, method)).HandlerFunc(h)
        }

        handleMethod("Create", createVolume)
        handleMethod("Remove", removeVolume)
        handleMethod("Mount", mountVolume)
        handleMethod("Unmount", unmountVolume)
        handleMethod("Path", volumePath)

        log.Fatal(http.ListenAndServe(":9080", router))
}

func notFound(w http.ResponseWriter, r *http.Request) {
        log.Println("[plugin] Not found: %+v", r)
        http.NotFound(w, r)
}

// Protocol Handlers
type handshakeResp struct {
        Implements []string
}

type errResp struct {
        Err string
}

type mountResp struct {
        Mountpoint string
        Err string
}

func handshake(w http.ResponseWriter, r *http.Request) {
        err := json.NewEncoder(w).Encode(&handshakeResp{
                []string{"VolumeDriver"},
        })
        if err != nil {
                log.Fatal("handshake encode:", err)
                sendError(w, "encode error", http.StatusInternalServerError)
                return
        }
        log.Println("Handshake complete!")
}

func createVolume(w http.ResponseWriter, r *http.Request) {
        err := json.NewEncoder(w).Encode(&errResp{
                "",
        })
        if err != nil {
                log.Fatal("createVolume encode:", err)
                sendError(w, "encode error", http.StatusInternalServerError)
                return
        }
        log.Println("createVolume complete!")
}

func removeVolume(w http.ResponseWriter, r *http.Request) {
        err := json.NewEncoder(w).Encode(&errResp{
                "",
        })
        if err != nil {
                log.Fatal("removeVolume encode:", err)
                sendError(w, "encode error", http.StatusInternalServerError)
                return
        }
        log.Println("removeVolume complete!")
}

func mountVolume(w http.ResponseWriter, r *http.Request) {
        err := json.NewEncoder(w).Encode(&mountResp{
                "/tmp/testvolume",
                "",
        })
        if err != nil {
                log.Fatal("mountVolume encode:", err)
                sendError(w, "encode error", http.StatusInternalServerError)
                return
        }
        log.Println("mountVolume complete!")
}

func unmountVolume(w http.ResponseWriter, r *http.Request) {
        err := json.NewEncoder(w).Encode(&errResp{
                "",
        })
        if err != nil {
                log.Fatal("unmountVolume encode:", err)
                sendError(w, "encode error", http.StatusInternalServerError)
                return
        }
        log.Println("unmountVolume complete!")
}

func volumePath(w http.ResponseWriter, r *http.Request) {
        err := json.NewEncoder(w).Encode(&mountResp{
                "/tmp/testvolume",
                "",
        })
        if err != nil {
                log.Fatal("volumePath encode:", err)
                sendError(w, "encode error", http.StatusInternalServerError)
                return
        }
        log.Println("volumePath complete!")
}

func sendError(w http.ResponseWriter, msg string, code int) {
        log.Print("%d %s", code, msg)
        http.Error(w, msg, code)
}
