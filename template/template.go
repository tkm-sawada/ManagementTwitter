package template

import (
	"ManagementTwitter/config"
	"ManagementTwitter/tweetapi"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

//htmlをキャッシング
var tpls = template.Must(template.ParseFiles("top.html", "search.html", "trend.html", "tweet.html", "reservation.html"))

func renderTemplate(w http.ResponseWriter, tpl string, data interface{}) {
	// t, _ := template.ParseFiles(tpl + ".html")
	// t.Execute(w, p)
	//キャッシングしたテンプレートに渡す
	err := tpls.ExecuteTemplate(w, tpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

///top/にアクセスされたときの処理
//　w http.ResponseWriter：リクエストに対するレスポンス
//　r *http.Request：リクエスト内容（ブラウザからの要求）
func topHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "top", nil)

}

///serch/にアクセスされたときの処理
//　w http.ResponseWriter：リクエストに対するレスポンス
//　r *http.Request：リクエスト内容（ブラウザからの要求）
func serchHandler(w http.ResponseWriter, r *http.Request) {

	//Path[len(string):] ・・・URLの[]の文字以降のアドレスを取得
	//title := r.URL.Path[len("/view/"):]
	word := r.FormValue("serchtext") //Formタグのnameを指定
	if len(word) == 0 {
		renderTemplate(w, "top", nil)
	} else {
		data := tweetapi.GetTweetSearch(word)
		//レスポンスに中身を格納
		//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
		renderTemplate(w, "search", data)
	}
}

///trend/にアクセスされたときの処理
//　w http.ResponseWriter：リクエストに対するレスポンス
//　r *http.Request：リクエスト内容（ブラウザからの要求）
func trendHandler(w http.ResponseWriter, r *http.Request) {

	data := tweetapi.GetTweetTrend()
	//レスポンスに中身を格納
	renderTemplate(w, "trend", data)
}

func tweetHandler(w http.ResponseWriter, r *http.Request) {

	word := r.FormValue("tweettext") //Formタグのnameを指定
	if len(word) == 0 {
		renderTemplate(w, "top", nil)
	} else {
		data := tweetapi.PostTweet(word)
		renderTemplate(w, "tweet", data)
	}

}

func reservationHandler(w http.ResponseWriter, r *http.Request) {

	var db config.Database

	v := r.FormValue("reservation")
	if v != "" {
		d := r.FormValue("date")
		d += ","
		d += r.FormValue("text")
		d += ","
		d += r.FormValue("flg")
		if config.Savefile([]byte(d)) != nil {
			return
		}
	}

	data, err := config.Readfile()
	if err != nil {
		renderTemplate(w, "top", nil)
	} else {
		arData := strings.Split(string(data), ",")

		db.Date = arData[0]
		db.Text = arData[1]
		db.Flg, _ = strconv.Atoi(arData[2])
		renderTemplate(w, "reservation", db)
	}

}
func HandlerCall() {
	http.HandleFunc("/", topHandler)
	http.HandleFunc("/serch/", serchHandler)
	http.HandleFunc("/trend/", trendHandler)
	http.HandleFunc("/tweet/", tweetHandler)
	http.HandleFunc("/reservation/", reservationHandler)
	http.ListenAndServe(":8080", nil)
}
