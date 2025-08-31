package GoPing

import (
	"errors"
	"math/rand"
	"net"
	"time"
)

// Lookup 根据给定的主机名查找对应的IP地址，并随机返回一个IP地址。
// 如果找不到对应的IP地址，则返回一个错误。
//
// 参数:
//
//	host - 需要查找IP地址的主机名。
//
// 返回值:
//
//	string - 随机选择的主机IP地址。
//	error - 如果查找失败，则返回错误信息。
func Lookup(host string) (string, error) {
	// 使用net包的LookupHost函数查询给定主机名的所有IP地址。
	adders, err := net.LookupHost(host)
	// 如果查询过程中发生错误，则直接返回空字符串和错误。
	if err != nil {
		return "", err
	}
	// 如果查询结果为空，则说明找不到对应的IP地址，返回一个自定义的错误。
	if len(adders) < 1 {
		return "", errors.New("unknown host")
	}
	// 使用随机数生成器来从查询到的IP地址列表中随机选择一个IP地址。
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 随机选择一个IP地址并返回。
	return adders[rd.Intn(len(adders))], nil
}

// Data 是一个全球唯一的，长度为32的字节切片。
// 它用于存储预定义的字符序列，该序列由小写字母组成，不包含特殊字符或数字。
// 这个预定义的序列在程序中可能被用作标识符、加密种子或其他需要固定序列的场景。
var Data = []byte("abcdefghijklmnopqrstuvwabcdefghi")

// Dail 方法用于建立到指定地址的ICMP连接。
// 它尝试通过IP协议族中的ICMP协议类型建立连接，如果失败，则返回错误。
//
// 参数:
//
//	pg *PingSet - 包含待连接地址信息的PingSet结构体指针。
//
// 返回值:
//
//	error - 如果建立连接过程中出现错误，则返回该错误；否则返回nil。
func (pg *PingSet) Dail() (err error) {
	// 尝试使用IP4协议和ICMP协议类型建立到pg.Addr指定地址的连接
	pg.Conn, err = net.Dial("ip4:icmp", pg.Addr)
	// 如果建立连接过程中出现错误，则直接返回该错误
	if err != nil {
		return err
	}
	// 如果连接成功建立，则返回nil表示无错误发生
	return nil
}

// SetDeadline 为PingSet对象的连接设置超时 deadline。
// 参数timeout表示超时时间，单位为秒。
// 返回值为操作可能产生的错误。
func (pg *PingSet) SetDeadline() error {
	// 根据当前时间计算出超时时间点，并为连接设置超时
	return pg.Conn.SetDeadline(time.Now().Add(pg.Timeout))
}

// Close 释放PingSet对象所持有的网络连接资源。
//
// 该方法通过调用PingSet对象内部Conn字段的Close方法来实现资源的释放。
// PingSet在使用完毕后，应该调用此方法来确保网络连接被正确关闭，避免资源泄露。
//
// 返回值:
//
//	该方法将Conn.Close()的返回值直接返回，因此调用者可以根据返回的error来判断关闭操作是否成功。
func (pg *PingSet) Close() error {
	return pg.Conn.Close()
}
