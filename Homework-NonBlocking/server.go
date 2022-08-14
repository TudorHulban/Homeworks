package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	engine *fiber.App
	cache  *cache
}

func (s *server) ProcSync(c *fiber.Ctx) error {
	duration := c.Params("duration")
	dur, errCo := strconv.Atoi(duration)
	if errCo != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide process duration")
	}

	if dur <= 0 {
		return c.Status(http.StatusBadRequest).SendString("please provide positive process duration")
	}

	procID := generatorProcID()
	newProcess(cfgProc{
		id:          procID,
		etaMilisecs: dur,
		c:           s.cache,
	})

	return c.SendString(fmt.Sprintf("Cache ID: %s\n", procID))
}

func (s *server) ProcASync(c *fiber.Ctx) error {
	duration := c.Params("duration")
	dur, errCo := strconv.Atoi(duration)
	if errCo != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide process duration: " + errCo.Error())
	}

	if dur <= 0 {
		return c.Status(http.StatusBadRequest).SendString("please provide positive process duration")
	}

	procID := generatorProcID()
	go newProcess(cfgProc{
		id:          procID,
		etaMilisecs: dur,
		c:           s.cache,
	})

	return c.SendString(fmt.Sprintf("Cache ID: %s\n", procID))
}

func (s *server) GetIDFromCache(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == strings.ToLower("all") {
		resp, errGet := s.cache.getAllProcessIDs()
		if errGet != nil {
			c.Status(http.StatusOK).SendString(fmt.Sprintf("%s\n", errGet))
			return nil
		}

		return c.SendString(fmt.Sprintf("%s\n", strings.Join(resp, "\n")))
	}

	p, errCa := s.cache.getProcessState(processID(id))
	if errCa != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("%s\n", errCa))
	}

	return c.SendString(fmt.Sprintf("%s\n", p.output))
}

func (s *server) GetRegister(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("%d\n", s.cache.getRegister()))
}

func (s *server) ManyProcSync(c *fiber.Ctx) error {
	duration := c.Params("duration")
	dur, errCo := strconv.Atoi(duration)
	if errCo != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide process duration:" + errCo.Error())
	}

	if dur <= 0 {
		return c.Status(http.StatusBadRequest).SendString("please provide positive process duration")
	}

	howMany := c.Params("howmany")
	many, errCo := strconv.Atoi(howMany)
	if errCo != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide how many processes" + errCo.Error())
	}

	if many <= 0 {
		return c.Status(http.StatusBadRequest).SendString("please provide positive number of processes")
	}

	var wg sync.WaitGroup
	var res []string

	for i := 0; i < many; i++ {
		procID := generatorProcID()
		wg.Add(1)

		go func() {
			defer wg.Done()
			newProcess(cfgProc{
				id:          procID,
				etaMilisecs: dur,
				c:           s.cache,
			})
		}()

		res = append(res, string(procID))
	}

	wg.Wait()

	return c.SendString(fmt.Sprintf("%s\n", strings.Join(res, "\n")))
}

func (s *server) ManyProcASync(c *fiber.Ctx) error {
	duration := c.Params("duration")
	dur, errCo := strconv.Atoi(duration)
	if errCo != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide process duration: " + errCo.Error())
	}

	if dur <= 0 {
		return c.Status(http.StatusBadRequest).SendString("please provide positive process duration")
	}

	howMany := c.Params("howmany")
	many, errCo := strconv.Atoi(howMany)
	if errCo != nil {
		return c.Status(http.StatusBadRequest).SendString("please provide how many processes: " + errCo.Error())
	}

	if many <= 0 {
		return c.Status(http.StatusBadRequest).SendString("please provide positive number of processes")
	}

	var res []string

	for i := 0; i < many; i++ {
		procID := generatorProcID()

		go newProcess(cfgProc{
			id:          procID,
			etaMilisecs: dur,
			c:           s.cache,
		})

		res = append(res, string(procID))
	}

	return c.SendString(fmt.Sprintf("%s\n", strings.Join(res, "\n")))
}
