package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"jeedev-api/myredis"
	"jeedev-api/units"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Area struct {
	Id    int64  `orm:"auto"`
	Name  string `orm:"size(128)"`
	Pid   int64
	Level int64
}

func init() {
	orm.RegisterModel(new(Area))
}

// AddArea insert a new Area into database and returns
// last inserted Id on success.
func AddArea(m *Area) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAreaById retrieves Area by Id. Returns error if
// Id doesn't exist
func GetAreaById(id int64) (v *Area, err error) {
	rs := myredis.Conn()
	key := "GetAreaById"+ strconv.FormatInt(id,10)
	// json数据在go中是[]byte类型，所以此处用redis.Bytes转换
	value, _ := redis.Bytes(rs.Do("GET", key))
	if value != nil {
		// 将json解析成map类型
		//v = &Area{}
		errShal := json.Unmarshal(value,&v)
		if errShal != nil {
			fmt.Println(errShal)
		}
		fmt.Println("from rediss")
		return v,nil
	}else {
		o := orm.NewOrm()
		v = &Area{Id: id}
		if err = o.QueryTable(new(Area)).Filter("Id", id).RelatedSel().One(v); err == nil {
			//存到redis
			value, _ := json.Marshal(v)  //转成json格式存起来
			_, err := rs.Do("SET", key, value)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("from db")

			return v, nil
		}
	}
	return nil, err
}

// GetAllArea retrieves all Area matches certain condition. Returns empty list if
// no records exist
func GetAllArea(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	rs := myredis.Conn()
	/*
	key := "GetAllArea"+units.Map2String(query) + units.Array2String(fields) + units.Array2String(sortby) + units.Array2String(order) + strconv.FormatInt(offset,10) + ":" + strconv.FormatInt(limit,10)
	fmt.Println(key)
	key = units.GetMd5(key)
	*/
	key := units.GetKey("GetAllArea",query,fields,sortby,order,offset,limit)
	fmt.Println(key)
	value,_ := redis.Bytes(rs.Do("GET",key))
	if value != nil {
		//将 value bytes转为 interface
		ml := units.Bytes2Intaface(value)
		fmt.Println("from redis")
		return ml, nil
	}else {
		o := orm.NewOrm()
		qs := o.QueryTable(new(Area))
		// query k=v
		for k, v := range query {
			// rewrite dot-notation to Object__Attribute
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}
		// order by:
		var sortFields []string
		if len(sortby) != 0 {
			if len(sortby) == len(order) {
				// 1) for each sort field, there is an associated order
				for i, v := range sortby {
					orderby := ""
					if order[i] == "desc" {
						orderby = "-" + v
					} else if order[i] == "asc" {
						orderby = v
					} else {
						return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
					}
					sortFields = append(sortFields, orderby)
				}
				qs = qs.OrderBy(sortFields...)
			} else if len(sortby) != len(order) && len(order) == 1 {
				// 2) there is exactly one order, all the sorted fields will be sorted by this order
				for _, v := range sortby {
					orderby := ""
					if order[0] == "desc" {
						orderby = "-" + v
					} else if order[0] == "asc" {
						orderby = v
					} else {
						return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
					}
					sortFields = append(sortFields, orderby)
				}
			} else if len(sortby) != len(order) && len(order) != 1 {
				return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
			}
		} else {
			if len(order) != 0 {
				return nil, errors.New("Error: unused 'order' fields")
			}
		}

		var l []Area
		qs = qs.OrderBy(sortFields...).RelatedSel()
		if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
			if len(fields) == 0 {
				for _, v := range l {
					ml = append(ml, v)
				}
			} else {
				// trim unused fields
				for _, v := range l {
					m := make(map[string]interface{})
					val := reflect.ValueOf(v)
					for _, fname := range fields {
						m[fname] = val.FieldByName(fname).Interface()
					}
					ml = append(ml, m)
				}
			}

			//存到redis
			value, _ := json.Marshal(ml)  //转成json格式存起来
			_, err := rs.Do("SET", key, value,"EX","300")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("from db")
			return ml, nil
		}
	}
	return nil, err
}

// UpdateArea updates Area by Id and returns error if
// the record to be updated doesn't exist
func UpdateAreaById(m *Area) (err error) {
	o := orm.NewOrm()
	v := Area{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArea deletes Area by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArea(id int64) (err error) {
	o := orm.NewOrm()
	v := Area{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Area{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
