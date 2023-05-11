package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Options struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type Response struct {
	Win  string `json:"win"`
	Lose string `json:"lose"`
}

type Quiz struct {
	Answer   uint      `json:"answer"`
	Question string    `json:"question"`
	Options  []Options `json:"options"`
	Response Response  `json:"response"`
}

func SendMockQuizzesToClient(c *fiber.Ctx) error {
	quizzes := []Quiz{
		{
			Answer:   1,
			Question: "این بار خیلی سخت بهت نمیگیرم؛ چون میترسم مغزت خیلی دووم نیاره! اگه گفتی صاحب شرکت آمازون اسمش چیه؟",
			Options: []Options{
				{ID: 1, Title: "جف بزوس"},
				{ID: 2, Title: "بیل گیتس"},
				{ID: 3, Title: "وارن بافت"},
				{ID: 4, Title: "مک کنیز بزوس"},
			},
			Response: Response{
				Win:  "ایول!",
				Lose: "جدا؟!",
			},
		},
		{
			Answer:   1,
			Question: "اینو هم آسون میگیرم که یه وقت فکر نکنی بی رحمم یا هرچی؛ اگه گفتی کدوم زبان بود که سال ۲۰۲۱ بهترین زبان برنامه نویسی جهان شد؟",
			Options: []Options{
				{ID: 1, Title: "جاوااسکریپت"},
				{ID: 2, Title: "پایتون"},
				{ID: 3, Title: "تایپ اسکریپت"},
				{ID: 4, Title: "راست Rust"},
			},
			Response: Response{
				Win:  "بابا خفن!",
				Lose: "نه عامو!",
			},
		},
		{
			Answer:   2,
			Question: "اگه گفتی کدوم انیمیشنه که بعد از ۳۰ فصل و بیشتر از ۳۰ سال هنوزم داره پخش میشه؟",
			Options: []Options{
				{ID: 1, Title: "انیمیشن پهلوانان"},
				{ID: 2, Title: "انیمیشن سیمپسون ها"},
				{ID: 3, Title: "انیمشین بهترین دوست"},
				{ID: 4, Title: "انیمیشن پدر آمریکایی"},
			},
			Response: Response{
				Win:  "El Barto!",
				Lose: "این همه ادعا؛ همین؟!",
			},
		},
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"quizzes": quizzes,
			"counts":  len(quizzes),
		},
		"status": "OK",
	})
}
