package main

import (
	"net"
	"fmt"
	"sync"
	"strings"
)

var wg	 sync.WaitGroup
const (
	Tcp4  string=  "tcp4"
	Tcp6  = "tcp6"
)

func IsErro(err error) bool  {
	return  err != nil
}

func main()  {
	tcpAddr,erro:= net.ResolveTCPAddr(Tcp4,":10099")
	if IsErro(erro) {
		fmt.Println(erro)
		return
	}
	fmt.Println(tcpAddr)
	LtnTcp,LtnErro := net.ListenTCP(Tcp4,tcpAddr)
	if IsErro(LtnErro) {
		fmt.Println(LtnErro)
		return
	}
 	EndNet  := false
	wg.Add(1)
	fmt.Println("开始Listen")
	go func() {
		defer wg.Done()
		for {
		  	conn,connErro := LtnTcp.Accept()
			fmt.Println("客户端连接")
			if IsErro(connErro) {
				fmt.Println("conn = ",connErro)
				conn.Close()
				fmt.Println("%v",conn)
				continue
			}
			if EndNet {
				break
			}
			go HandlerRead(conn)
		}
		LtnTcp.Close()
	}()
	wg.Wait()
	fmt.Println("MainEnd")
}

func HandlerRead(conn net.Conn)  {
	for  {
		msgBuf := make([]byte,65535)
		buflen,erro := conn.Read(msgBuf)
		if IsErro(erro) {
			fmt.Println(erro)
			// 客户端主动断开连接
			conn.Close()
			return
		}
		recvstr := string(msgBuf)
		fmt.Println("接收数据长度",buflen,recvstr)
		if  strings.Compare(recvstr,"Q")  == 0 {
			fmt.Println("客户端断开连接",recvstr)
			conn.Close()
		}
	}

}


func HandlerWrite(conn net.Conn,msgBuf []byte)  {
	//buflen,erro := conn.Read(msgBuf)
	//if IsErro(erro) {
	//
	//}
}