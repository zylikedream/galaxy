/*
 * @Author: your name
 * @Date: 2021-10-19 17:41:17
 * @LastEditTime: 2021-11-04 17:05:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /components/gxynet/conn/endpoint.go
 */
package endpoint

import (
	"context"
	"net"
)

type Endpoint interface {
	Send(msg interface{}) error
	Start(ctx context.Context) error
	Close(context.Context, error)
	Conn() net.Conn
	GetData() interface{}
	SetData(interface{})
}
