package cached

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// ADD key value
// DEL key
// GET key
// UPDATE key value
// GETALL
// DELALL
//
// OK
// ERROR
// FAIL

type CacheDServer struct {
	cd *CacheD
}

func NewServer() *CacheDServer {
	return &CacheDServer{
		cd: NewCacheD(),
	}
}

func (s *CacheDServer) Start() {
	address := "localhost:8080"
	server, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("erro ao criar servidor em %s: %s\n", address, err)
		return
	}

	fmt.Printf("Iniciando servidor em %s\n", address)

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("erro ao abrir conexão com cliente: %s\n", err)
			continue
		}

		go s.process(conn)
	}
}

func (s *CacheDServer) process(conn net.Conn) {
	defer conn.Close()

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("falha ao ler dados da conexão:", err)
		fmt.Fprintf(conn, "ERROR\n")
		return
	}

	data = strings.TrimSpace(data)
	fmt.Printf("\x1b[1;36m%s\x1b[0m\n", data)
	command := strings.SplitN(data, " ", 3)
	invalidCommand := false

	switch len(command) {
	case 3:
		cmd, k, v := command[0], command[1], command[2]
		switch cmd {
		case "ADD":
			if err := s.cd.Add(k, v); err != nil {
				fmt.Printf("erro ao executar ADD: %s\n", err)
				fmt.Fprintf(conn, "FAIL\n")
			} else {
				fmt.Fprintf(conn, "OK\n")
			}

		case "UPDATE":
			s.cd.Update(k, v)
			fmt.Fprintf(conn, "OK\n")

		default:
			invalidCommand = true
		}

	case 2:
		cmd, k := command[0], command[1]
		switch cmd {
		case "GET":
			v, err := s.cd.Get(k)
			if err != nil {
				fmt.Printf("erro ao executar GET: %s\n", err)
				fmt.Fprintf(conn, "FAIL\n")
			} else {
				fmt.Fprintf(conn, "%s\nOK\n", v)
			}

		case "DEL":
			if err := s.cd.Del(k); err != nil {
				fmt.Printf("erro ao executar DEL: %s\n", err)
				fmt.Fprintf(conn, "FAIL\n")
			} else {
				fmt.Fprintf(conn, "OK\n")
			}

		default:
			invalidCommand = true
		}

	case 1:
		switch command[0] {
		case "GETALL":
			pairs := s.cd.GetAll()
			for _, pair := range pairs {
				fmt.Fprintf(conn, "%s\t%s\n", pair[0], pair[1])
			}
			fmt.Fprintf(conn, "OK\n")

		case "DELALL":
			s.cd.DelAll()
			fmt.Fprintf(conn, "OK\n")

		default:
			invalidCommand = true
		}

	default:
		fmt.Printf("comando não segue o contrato necessário: %v\n", command)
		fmt.Fprintf(conn, "FAIL\n")
		return
	}

	if invalidCommand {
		fmt.Printf("comando inválido: %s\n", data)
		fmt.Fprintf(conn, "ERROR\n")
	}
}
