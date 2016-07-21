package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Host struct {
	Id           int64
	Class        string `orm:size(50)`
	Service_name string `orm:size(50)`
	Name         string `orm:size(50)`
	Ip           string `orm:size(15)`
	Port         string `orm:size(15)`
	Os_type      string `orm:size(50)`
	Owner        string `orm:size(50)`
	Cpu          string `orm:size(50)`
	Mem          string `orm:size(50)`
	Disk         string `orm:size(50)`
	Idc          string `orm:size(50)`
	Root         string `orm:size(10)`
	Rootpwd      string `orm:size(50)`
	Read         string `orm:size(10)`
	Readpwd      string `orm:size(50)`
	Group        string `orm:size(50)`
	Comment      string `orm:size(100)`
	Created      time.Time
}

// register host model
func init() {
	orm.RegisterModel(new(Host))
}

func GetHostCount() (int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	total, err := o.QueryTable("host").All(&hosts)
	if err != nil {
		return 0, err
	}
	return total, err
}

func GetHosts(currPage, pageSize int) ([]*Host, int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	total, err := o.QueryTable("host").Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
	if err != nil {
		return nil, 0, err
	}
	return hosts, total, err
}

func GetHostById(id string) (*Host, error) {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{}
	err = o.QueryTable("host").Filter("id", hid).One(host)
	return host, err
}

func AddHost(class, service_name, name, ip, port, os_type, owner, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	rootpwd, _ = AESEncode(rootpwd, AesKey)
	readpwd, _ = AESEncode(readpwd, AesKey)
	host := &Host{
		Class:        class,
		Service_name: service_name,
		Name:         name,
		Ip:           ip,
		Port:         port,
		Os_type:      os_type,
		Owner:        owner,
		Root:         root,
		Read:         read,
		Rootpwd:      rootpwd,
		Readpwd:      readpwd,
		Cpu:          cpu,
		Mem:          mem,
		Disk:         disk,
		Group:        group,
		Idc:          idc,
		Comment:      comment,
		Created:      time.Now(),
	}
	err := o.QueryTable("host").Filter("name", name).Filter("service_name", service_name).Filter("ip", ip).One(host)
	if err == nil {
		msg = "主机名或服务名已存在"
		return nil, msg
	}
	_, err = o.Insert(host)
	msg = "添加主机成功"
	return err, msg
}

func ModifyHost(id, class, service_name, name, ip, port, os_type, owner, root, read, rootpwd, readpwd, cpu, mem, disk, group, idc, comment string) (error, string) {
	o := orm.NewOrm()
	var msg string
	rootpwd, _ = AESEncode(rootpwd, AesKey)
	readpwd, _ = AESEncode(readpwd, AesKey)
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{
		Id: hid,
	}
	err = o.Read(host)
	if err == nil {
		host.Class = class
		host.Service_name = service_name
		host.Name = name
		host.Ip = ip
		host.Port = port
		host.Os_type = os_type
		host.Owner = owner
		host.Root = root
		host.Read = read
		host.Rootpwd = rootpwd
		host.Readpwd = readpwd
		host.Cpu = cpu
		host.Mem = mem
		host.Disk = disk
		host.Group = group
		host.Idc = idc
		host.Comment = comment
	}
	o.Update(host)
	msg = "修改成功"
	return err, msg
}

func DeleteHost(id string) error {
	o := orm.NewOrm()
	hid, err := strconv.ParseInt(id, 10, 64)
	host := &Host{
		Id: hid,
	}
	_, err = o.Delete(host)
	if err != nil {
		return err
	}
	return nil
}

func SearchHostCount(idc, name string) (int64, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	total, err := o.QueryTable("host").Filter("idc", idc).Filter("name__icontains", name).All(&hosts)
	return total, err
}

type Name struct {
	name string
	ip   string
}

func SearchHostByName(currPage, pageSize int, idc, name string) ([]*Host, error) {
	o := orm.NewOrm()
	hosts := make([]*Host, 0)
	/*
		var cond *orm.Condition
		cond = orm.NewCondition()
		cond = cond.Or("name__icontains", name)
		cond = cond.Or("ip__icontains", "ip")
		var qs orm.QuerySeter
		qs = o.QueryTable("host").Filter("idc", idc).Limit(pageSize, (currPage-1)*pageSize).SetCond(cond)
		_, err := qs.All(&hosts)
	*/
	_, err := o.QueryTable("host").Filter("idc", idc).Filter("name__icontains", name).Limit(pageSize, (currPage-1)*pageSize).All(&hosts)
	return hosts, err
}
