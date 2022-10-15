package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type SimpleTelnetClient struct {
	conn    net.Conn
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
}

func (c *SimpleTelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return fmt.Errorf("SimpleTelnetClient: cannot connect to %s: %w", c.address, err)
	}

	c.conn = conn

	return nil
}

func (c *SimpleTelnetClient) Send() error {
	scanner := bufio.NewScanner(c.in)

	for scanner.Scan() {
		b := append(scanner.Bytes(), '\n')
		_, err := c.conn.Write(b)
		if err != nil {
			return fmt.Errorf("...Connection was closed by peer")
		}
	}

	return nil
}

func (c *SimpleTelnetClient) Receive() error {
	scanner := bufio.NewScanner(c.conn)

	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) > 0 {
			b = append(scanner.Bytes(), '\n')
			_, err := c.out.Write(b)
			if err != nil {
				return fmt.Errorf("SimpleTelnetClient: cannot write to output: %w", err)
			}
		}
	}

	return nil
}

func (c *SimpleTelnetClient) Close() error {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return fmt.Errorf("SimpleTelnetClient: cannot close connection to %s: %w", c.address, err)
		}
	}

	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &SimpleTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
