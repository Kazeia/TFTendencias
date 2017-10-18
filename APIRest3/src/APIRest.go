////DOCUMENTATION
//https://godoc.org/gopkg.in/mgo.v2
//http://www.jancarloviray.com/blog/go-mongodb-simple-example/
//PREGUNTAR LAS CREDENCIALES DE USUARIO DE BASE DE DATOS -> ./DB/Access.go
package main

//MEJORAR:
//  1. Pasar los parametros como direccion de memoria, para eliminar la sesion y variables
//  2. Por ahora estamos enviando desde ./Test/SendDotNet/Send#, esto es solo de prueba, xq debe ser enviado desde el ClienteWeb.
import (
	"log"
	"net/http"
	"sync"

	Routes "./Routes"

	"github.com/julienschmidt/httprouter"
	//	"github.com/streadway/amqp"
)

//Parallelize funcion personalizada para que pueda correr mas de 1 gorutina
func Parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

//**-MAIN-**
func main() {

	listenFirstChannel := func() {
		// conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
		// defer conn.Close()

		// ch, _ := conn.Channel()
		// defer ch.Close()

		// q, _ := ch.QueueDeclare(
		// 	"canal1", // name
		// 	false,    // durable
		// 	false,    // delete when unused
		// 	false,    // exclusive
		// 	false,    // no-wait
		// 	nil,      // arguments
		// )

		// msgs, _ := ch.Consume(
		// 	q.Name, // queue
		// 	"",     // consumer
		// 	true,   // auto-ack
		// 	false,  // exclusive
		// 	false,  // no-local
		// 	false,  // no-wait
		// 	nil,    // args
		// )

		// log.Printf("Esperando Mensajes en Canal1...")

		// for d := range msgs {
		// 	log.Printf("Received a message: %s", d.Body)
		// }
	}

	listenSecondChannel := func() {
		// conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
		// defer conn.Close()

		// ch, _ := conn.Channel()
		// defer ch.Close()

		// q, _ := ch.QueueDeclare(
		// 	"canal2", // name
		// 	false,    // durable
		// 	false,    // delete when unused
		// 	false,    // exclusive
		// 	false,    // no-wait
		// 	nil,      // arguments
		// )

		// msgs, _ := ch.Consume(
		// 	q.Name, // queue
		// 	"",     // consumer
		// 	true,   // auto-ack
		// 	false,  // exclusive
		// 	false,  // no-local
		// 	false,  // no-wait
		// 	nil,    // args
		// )

		// log.Printf("Esperando Mensajes en Canal2...")

		// for d := range msgs {
		// 	log.Printf("Received a message: %s", d.Body)
		// }
	}

	listenServer := func() {
		//creacion de enrutador pat
		router := httprouter.New()

		//Enrutadores-User
		router.GET("/users", Routes.GetAllUser)
		router.GET("/users/:id", Routes.GetUserByID)
		router.POST("/users", Routes.PostUser)
		router.PUT("/users/:id", Routes.PutUserByID)
		router.DELETE("/users/:id", Routes.DeleteUserByID)

		//Enrutadores-Flight
		router.GET("/flights", Routes.GetAllFlight)
		router.GET("/flights/:id", Routes.GetFlightByID)
		router.POST("/flights", Routes.PostFlight)
		router.PUT("/flights/:id", Routes.PutFlightByID)
		router.DELETE("/flights/:id", Routes.DeleteFlightByID)

		//Enrutadores-Booking
		router.GET("/bookings", Routes.GetAllBooking)
		router.GET("/bookings/:id", Routes.GetBookingByID)
		router.POST("/bookings", Routes.PostBooking)
		router.PUT("/bookings/:id", Routes.PutBookingByID)
		router.DELETE("/bookings/:id", Routes.DeleteBookingByID)

		//Escuchando el servidor
		log.Print("Escuchando en 127.0.0.1:8300...")
		http.ListenAndServe(":8300", router)
	}

	Parallelize(listenFirstChannel, listenSecondChannel, listenServer)
}
