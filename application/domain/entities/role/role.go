package entities

type Role struct {
	ID      string `firestore:"-"`
	Nome    string `firestore:"nome" binding:"required,min=5,max=50"`
	Local   string `firestore:"local" binding:"required,min=10,max=80"`
	Cidade  string `firestore:"cidade" binding:"required,min=10,max=50"`
	Horario string `firestore:"horario" binding:"required,min=5,max=15"`
	Dia     string `firestore:"dia" binding:"required,datetime=02-01-2006"`
	Tipo    string `firestore:"tipo_role" binding:"required,min=5,max=40"`
}
