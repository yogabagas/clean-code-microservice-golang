package controller

type AppController struct {
	Student interface{ StudentController }
}
