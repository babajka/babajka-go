package main

// data.go contains data to populate database on initialization.

var (
	articles = map[int]Article{
		1: Article{
			ID:       1,
			Title:    "Вальжына Морт",
			Author:   "Вальжына Морт",
			Hashtags: []string{"#паэзія", "#морт"},
			Photo:    "/images/photo8.jpg",
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
		2: Article{
			ID:       2,
			Title:    "Эмігранцкія гісторыі",
			Author:   "34mag.net",
			Hashtags: []string{"34magnet", "#34"},
			Photo:    "/images/34magnet.jpg",
			Text: `Я нарадзілася ў Горадні. І раней я заўсёды распавядала людзям пра тое, што мой горад вельмі прыгожы. Там утульныя спальныя раёны, шмат паркаў і чыстыя вуліцы. Выдатны старадаўні цэнтр. Распавядала, што ў гэтым горадзе калісьці жыў Вітаўт, які пабудаваў адзін з замкаў на беразе Нёмана. Распавядала пра тое, што Нёман раней зваўся Кронанам, таму што ў горадзе знаходзіліся каралеўскія рэзідэнцыі.
Я выдатна скончыла школу, у мяне было шмат сяброў і маса захапленняў накшталт стральбы, верхавой язды і танцаў, якія я сумяшчала з падрыхтоўкай да паступлення ў медунівер і сканчэннем музычнай школы. У мяне была каханая дзяўчына, у мяне былі шчаслівыя і здаровыя бацькі. Вялікая кватэра, машына, лецішча. Маё жыццё было ідэальным.`,
		},
		3: Article{
			ID:        3,
			Title:     "«Хулігану» – 4!",
			Author:    "34mag.net",
			Hashtags:  []string{"34magnet", "#34", "#кастрычніцкая"},
			Photo:     "/images/huligan.jpg",
			IsSpecial: true,
			Text:      `Дзве ночы запар наш улюбёны бар «Хуліган» душэўна адзначаў свой чацвёрты дзень нараджэння. Як заўсёды, тут сустракаліся старыя знаёмыя і заводзіліся новыя. Гучалі віншаванні. Ілля Чарапко какетліва перакідваўся слоўцамі з публікай, а By Soulfam разам з Пафнуціем спантанна зладзілі джэм на ганку барчэлы. Дзякуй вам за гэтыя выходныя. Long live «Хуліган»!`},
	}
)
