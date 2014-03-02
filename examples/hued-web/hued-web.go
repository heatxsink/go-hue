package main

import (
	"encoding/json"
	"net/http"
	"net/http/fcgi"
	"net"
	"fmt"
	"github.com/golang/glog"
	"../../src/lights"
	"../../src/portal"
	"../../src/groups"
	"../../src/key"
	"flag"
)

var (
	username_api_key_filename = ""
	username_api_key = ""
	hostname = ""
	hue_hostname = ""
	local_web_server_flag = true
	web_bind_address= "127.0.0.1"
	web_bind_port = 9000
)

type ApiResponse struct {
	Result bool `json:"result"`
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
}

func group_name_presets(name string) int {
	return_value := -1
	if name == "all" {
		return_value = 0
	} else if name == "office" {
		return_value = 1
	} else if name == "bedroom" {
		return_value = 2
	} else if name == "living-room" {
		return_value = 3
	} else if name == "upstairs" {
		return_value = 4
	}
	return return_value
}

func light_state_presets(name string) lights.State {
	on := lights.State { On: true }
	off := lights.State { On: false }
	red := lights.State { On: true, Hue: 65527, Effect: "none", Bri: 13, Sat: 253, Ct: 500, Xy: []float32{0.6736, 0.3221}, Alert: "none", TransitionTime: 4 }
	blue := lights.State { On: true, Hue: 46573, Effect: "none", Bri: 254, Sat: 251, Ct: 500, Xy: []float32{0.1754, 0.0556}, Alert: "none", TransitionTime: 4 }
	energize := lights.State { On: true, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, Ct: 155, Xy: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: 4 }
	relax := lights.State { On: true, Hue: 13088, Effect: "none", Bri: 144, Sat: 212, Ct: 467, Xy: []float32{0.5128, 0.4147}, Alert: "none", TransitionTime: 4 }
	reading := lights.State { On: true, Hue: 15331, Effect: "none", Bri: 222, Sat: 121, Ct: 343, Xy: []float32{0.4448, 0.4066}, Alert: "none", TransitionTime: 4 }
	concentrate := lights.State { On: true, Hue: 33849, Effect: "none", Bri: 219, Sat: 44, Ct: 234, Xy: []float32{0.3693, 0.3695}, Alert: "none", TransitionTime: 4 }
	candle_light := lights.State { On: true, Hue: 15339, Effect: "none", Bri: 19, Sat: 120, Ct: 343, Xy: []float32{0.4443, 0.4064}, Alert: "none", TransitionTime: 4 }
	virgin_america := lights.State { On: true, Hue: 54179, Effect: "none", Bri: 230, Sat: 253, Ct: 223, Xy: []float32{0.3621, 0.1491}, Alert: "none", TransitionTime: 4 }
	white := lights.State { On: true, Hue: 34495, Effect: "none", Bri: 203, Sat: 232, Ct: 155, Xy: []float32{0.3151, 0.3252}, Alert: "none", TransitionTime: 4 }
	orange := lights.State { On: true, Hue: 4868, Effect: "none", Bri: 254, Sat: 252, Ct: 500, Xy: []float32{0.6225, 0.3594}, Alert: "none", TransitionTime: 4 }
	deep_sea := lights.State { On: true, Hue: 65527, Effect: "none", Bri: 253, Sat: 253, Ct: 500, Xy: []float32{0.6736, 0.3221}, Alert: "none", TransitionTime: 4 }
	green := lights.State { On: true, Hue: 25654, Effect: "none", Bri: 254, Sat: 253, Ct: 290, Xy: []float32{0.4083, 0.5162}, Alert: "none", TransitionTime: 4 }
	snow := lights.State { On: true, Hue: 34258, Effect: "none", Bri: 254, Sat: 176, Ct: 181, Xy: []float32{0.3327, 0.3413}, Alert: "none", TransitionTime: 4 }
	movie_mode := lights.State { On: true, Hue: 65527, Effect: "none", Bri: 51, Sat: 253, Ct: 500, Xy: []float32{0.6736, 0.3221}, Alert: "none", TransitionTime: 4 }
	// default ...
	return_value := on
	// on with the show!
	if name == "on" {
		return on
	} else if name == "off" {
		return off
	} else if name == "red" {
		return red
	} else if name == "blue" {
		return blue
	} else if name == "energize" {
		return energize
	} else if name == "relax" {
		return relax
	} else if name == "reading" {
		return reading
	} else if name == "concentrate" {
		return concentrate
	} else if name == "candle-light" {
		return candle_light
	} else if name == "virgin-america" {
		return virgin_america
	} else if name == "white" {
		return white
	} else if name == "orange" {
		return orange
	} else if name == "deep-sea" {
		return deep_sea
	} else if name == "green" {
		return green
	} else if name == "snow" {
		return snow
	} else if name == "movie-mode" {
		return movie_mode
	}
	return return_value
}

func groupV1(w http.ResponseWriter, req *http.Request) {
	api_response := ApiResponse{ Result: true, Message: "", StatusCode: http.StatusOK }
	req.ParseForm()
	query_params := req.Form
	if req.Method == "GET" {
		name, name_exists := query_params["name"]
		state, state_exists := query_params["state"]

		if name_exists {
			group_id := group_name_presets(name[0])
			if state_exists {
				gg := groups.NewGroup(hue_hostname, username_api_key)
				group_state := light_state_presets(state[0])
				gg.SetGroupState(group_id, group_state)
			} else {
				api_response.Result = false
				api_response.Message = "Invalid state."
				api_response.StatusCode = http.StatusUnauthorized
			}
		} else {
			api_response.Result = false
			api_response.Message = "Invalid id or name."
			api_response.StatusCode = http.StatusUnauthorized
		}
	} else {
		api_response.Result = false
		api_response.Message = "Not an HTTP GET."
		api_response.StatusCode = http.StatusForbidden
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(api_response.StatusCode)
	json_data, _ := json.Marshal(&api_response)
	w.Write([]byte(json_data))
}

func statusV1(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte(http.StatusText(http.StatusTeapot)))
}

func user_interface_phoneV1(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./www/phone.html")
}

func user_interface_tabletV1(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./www/tablet.html")
}

func static_root(w http.ResponseWriter, req *http.Request) {
	if len(req.URL.Path) == 1 {
		root(w, req)
	} else if string(req.URL.Path[1:7]) == "static" {
		path := fmt.Sprintf("./www/%s", req.URL.Path[1:])
		http.ServeFile(w, req, path)
	} else {
		root(w, req)
	}
}

func root(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(http.StatusText(http.StatusNotFound)))
}

func check_error(err error) {
	if err != nil {
		glog.Fatalf("Error: %s\n", err.Error())
		glog.Flush()
	}
}

func web_start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/1/group", groupV1)
	mux.HandleFunc("/api/1/status", statusV1)
	mux.HandleFunc("/tablet", user_interface_tabletV1)
	mux.HandleFunc("/phone", user_interface_phoneV1)
	mux.HandleFunc("/", static_root)
	if local_web_server_flag {
		hostname := fmt.Sprintf(":%d", web_bind_port)
		start_message := fmt.Sprintf("Starting local hued-web on %s\n", hostname)
		fmt.Println(start_message)
		glog.Infof(start_message)
		http.ListenAndServe(hostname, mux)
	} else {
		hostname := fmt.Sprintf("%s:%d", web_bind_address, web_bind_port)
		l, err := net.Listen("tcp", hostname)
		if err == nil {
			start_message := fmt.Sprintf("Starting hued-web on %s\n", hostname)
			fmt.Println(start_message)
			glog.Infof(start_message)
			fcgi.Serve(l, mux)
		} else {
			glog.Error("Error starting server, is the port already in use?")
		}
	}
}

func init() {
	flag.StringVar(&username_api_key_filename, "conf", "", "The conf file that has the hue username api key in it.")
	flag.Parse()
}

func main() {
	k := key.New(username_api_key_filename)
	username_api_key = k.Username
	portal := portal.GetPortal()
	hue_hostname = portal[0].InternalIpAddress
	web_start()
}
