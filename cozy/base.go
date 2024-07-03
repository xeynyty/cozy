package cozy

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Core struct {
	config map[string]any
}

func (c *Core) Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return err
			}
		}

		if len(bytes) == 0 {
			break
		}

		err = c.getData(string(bytes))
		if err != nil {
			return err
		}
	}

	// It's like a defer file.Close()
	return file.Close()
}

func (c *Core) Get(key string, out *any) error {
	if val, ok := c.config[key]; ok {
		*out = val
		return nil
	}
	return errors.New(fmt.Sprintf("key '%v' not found in config", key))
}

func (c *Core) Set(key string, value any) {
	c.config[key] = value
}

func (c *Core) getData(line string) error {

	if strings.HasPrefix(line, "#") {
		return nil
	}

	keyValue := strings.Split(line, "=")
	if len(keyValue) != 2 {
		return errors.New(fmt.Sprintf("incorrect line in config: '%v'", line))
	}

	key := strings.TrimSpace(keyValue[0])
	value := strings.TrimSpace(keyValue[1])

	c.config[key] = value

	return nil
}

func (c *Core) Print() {
	fmt.Println(c.config)
}

var instance *Core
var once = &sync.Once{}

func GetInstance() *Core {
	once.Do(func() {
		instance = &Core{
			config: make(map[string]any),
		}
	})
	return instance
}
