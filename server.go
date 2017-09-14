package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func handleArticle(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	article := articles[1]
	aid, _ := strconv.Atoi(ss[2])
	article.ID = aid
	fmt.Println("requesting article", aid)
	tmpl, _ := template.ParseFiles("./templates/article.html")
	tmpl.Execute(w, article)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/login.html")
	tmpl.Execute(w, articles[1])
}

func defineStaticRoutes() {
	for _, r := range []string{"/styles/", "/images/", "/fonts/"} {
		http.Handle(r, http.StripPrefix(r, http.FileServer(http.Dir("."+r))))
	}
}

func initDB() {
	_, err := sql.Open("postgres", "user=uladbohdan dbname=wir-prod sslmode=verify-full")
	if err != nil {
		fmt.Println("DB ERR", err)
		return
	}

}

func main() {
	fmt.Println("initializing server...")

	http.Handle("/", http.FileServer(http.Dir("./templates")))

	defineStaticRoutes()

	http.HandleFunc("/articles/", handleArticle)
	http.HandleFunc("/login/", handleLogin)

	port := os.Getenv("PORT")

	fmt.Printf("listening on port :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

// Article ololo.
type Article struct {
	ID       int
	Title    string
	Text     string
	Author   string
	Hashtags []string
}

// ArticlePage comment.
type ArticlePage struct {
	Article  Article
	Featured []int
}

var (
	articles = map[int]Article{
		1: Article{
			ID:    1,
			Title: "Вальжына Морт",
			Text: `– Вы доўгі час жывяце не ў Беларусі. Як асэнсоўваеце краіну ў новых творах?
      – Для мяне Беларусь пашыраецца не ўшыр, а ўглыб. У маёй новай кнізе вершы пабудаваныя вакол вобраза
      дрэва,
      якое расце з крыві, касцей, чалавечыны і нашай зямлі. Дрэва расце ўгару, размаўляе з намі лісцем шэптам,
      які
      нарадзіўся раней за вусны. Мне цікавы шэпт нашых мёртвых, нашых прывідаў. Гэта Беларусь невербалізаваных
      жыццяў, якія не хочуць знікнуць і маўчаць, каб мы іх пачулі.
      – Што здарылася з вашымі кнігамі ў перыяд з 2005 па 2017?
      – Паміж 2005 і 2017 у мяне выйшлі кніжкі ў ЗША, Нямеччыне і Швецыі. У «Эпідэміі ружаў» ёсць вершы з
      гэтых
      кніжак, і трэць з іх – новыя. У першую чаргу паэт – гэта вушы і вочы. Я ўважліва сачу за прывідамі,
      слухаю
      сваіх мёртвых. Яны жывуць сярод нас. Іх прысутнасць можна пабачыць на нашых тварах.
      – Мёртвыя – гэта продкі або гэта нейкая метафара?
      – Я не бачу розніцы паміж рэальнасцю і метафарай. Літаратура змешвае памяць і ўяўленне так, што
      немагчыма
      адрозніць адно ад аднаго. У нашай прыродзе так і ёсць: памяць не можа існаваць без уяўлення, а ўяўленне
      не
      можа існаваць без памяці. Гэтыя два паняцці існуюць разам.
      Разам яны шукаюць нейкую форму: літаратурную, музычную, кінематаграфічную. Тут няма розніцы. Наогул
      паэзію і
      літаратуру я бачу як навуку. Біялогія чалавека, анатомія чалавека і літаратурная метафара існуюць для
      мяне ў
      адной прасторы.
      Метафара – гэта не гуманітарная абстракцыя, а элемент навуковага падыходу да свету. Гэта жаданне знайсці
      матэматычную або фізічную формулу, з дапамогай якой мы скажам: «Я – гэта ты», «Ты – гэта кніга», «Кніга
      –
      гэта паліца кніг», «Паліца – гэта дрэва, з якой яна зроблена, а дрэва – гэта я». Мы ўсе звязаныя адной
      формулай: розныя рэчы ствараюць асобную прастору, дзе яны могуць быць адным, застаючыся пры гэтым сабою.`,
		},
	}
)
