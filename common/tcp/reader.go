package tcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"time"
)

func ReadDataWithTimeout(conn *net.TCPConn, buf []byte, timeout time.Duration) error {
	err := (*conn).SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return err
	}

	c, err := (*conn).Read(buf)
	if err != nil {
		var opErr *net.OpError
		if errors.As(err, &opErr) && opErr.Timeout() {
			// 抛出自定义异常，表示读取超时
			return fmt.Errorf("read timeout")
		}
		return err
	}

	if c == 0 {
		return fmt.Errorf("connection closed")
	}

	return nil
}

func ReadData(conn *net.TCPConn) ([]byte, error) {
	var dataLen uint32
	dataLenBuf := make([]byte, 4)
	if err := readFixedData(conn, dataLenBuf); err != nil {
		return nil, err
	}
	// fmt.Printf("readFixedData:%+v\n", dataLenBuf)
	buffer := bytes.NewBuffer(dataLenBuf)
	if err := binary.Read(buffer, binary.BigEndian, &dataLen); err != nil {
		return nil, fmt.Errorf("read headlen error:%s", err.Error())
	}
	if dataLen <= 0 {
		return nil, fmt.Errorf("wrong headlen :%d", dataLen)
	}
	dataBuf := make([]byte, dataLen)
	// fmt.Printf("readFixedData.dataLen:%+v\n", dataLen)
	if err := readFixedData(conn, dataBuf); err != nil {
		return nil, fmt.Errorf("read headlen error:%s", err.Error())
	}
	return dataBuf, nil
}

// 读取固定buf长度的数据
func readFixedData(conn *net.TCPConn, buf []byte) error {
	_ = (*conn).SetReadDeadline(time.Now().Add(time.Duration(120) * time.Second))
	var pos int = 0
	var totalSize int = len(buf)
	for {
		c, err := (*conn).Read(buf[pos:])
		if err != nil {
			return err
		}
		pos = pos + c
		if pos == totalSize {
			break
		}
	}
	return nil
}
