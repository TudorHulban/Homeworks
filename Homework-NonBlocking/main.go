package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	s := server{
		engine: fiber.New(),
		cache:  newCache(),
	}

	s.engine.Get("one/sync/:duration", s.ProcSync)
	s.engine.Get("one/async/:duration", s.ProcASync)

	s.engine.Get("many/:howmany/sync/:duration", s.ManyProcSync)
	s.engine.Get("many/:howmany/async/:duration", s.ManyProcASync)

	s.engine.Get("cache/:id", s.GetIDFromCache)
	s.engine.Get("reg", s.GetRegister)

	log.Fatal(s.engine.Listen(":8080"))
}
