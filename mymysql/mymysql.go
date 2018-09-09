package mymysql

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"jeedev-api/myredis"
	"jeedev-api/units"
	"reflect"
	"strings"
)

var table interface{}

func TableName(T interface{}) {
	table = T
}

//query map[string]string, fields []string, sortby []string, order []string,
//	offset int64, limit int64
func FindAll(params map[string]interface{}, i interface{}) (ml []interface{}, err error) {
	//var query map[string]string
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if params["fields"] != nil {
		if v := params["fields"].(string); v != "" {
			fields = strings.Split(v, ",")
		}
	}

	// limit: 10 (default is 10)
	if params["limit"] != nil {
		if v := params["limit"].(int64); v != 0 {
			limit = v
		}
	}
	// offset: 0 (default is 0)
	if params["offset"] != nil {
		if v := params["offset"].(int64); v != 0 {
			offset = v
		}
	}

	// sortby: col1,col2
	if params["sortby"] != nil {
		if v := params["sortby"].(string); v != "" {
			sortby = strings.Split(v, ",")
		}
	}
	// order: desc,asc
	if params["order"] != nil {
		if v := params["order"].(string); v != "" {
			order = strings.Split(v, ",")
		}
	}
	// query: k:v,k:v
	if params["query"] != nil {

		if q := params["query"].(string); q != "" {
			for _, cond := range strings.Split(q, ",") {
				kv := strings.SplitN(cond, ":", 2)
				if len(kv) != 2 {
					errors.New("Error: invalid query key/value pair")
					return
				}
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}
	}
	rs := myredis.Conn()
	/*
		key := "GetAllArea"+units.Map2String(query) + units.Array2String(fields) + units.Array2String(sortby) + units.Array2String(order) + strconv.FormatInt(offset,10) + ":" + strconv.FormatInt(limit,10)
		fmt.Println(key)
		key = units.GetMd5(key)
	*/
	s := fmt.Sprintf("%s", reflect.TypeOf(table))
	fmt.Println(s)
	key := units.GetKey("GetAll"+s, query, fields, sortby, order, offset, limit)
	//fmt.Println(key)
	value, _ := redis.Bytes(rs.Do("GET", key))
	if value != nil {
		//将 value bytes转为 interface
		ml := units.Bytes2Intaface(value)
		fmt.Println("from redis")
		return ml, nil
	} else {
		o := orm.NewOrm()
		qs := o.QueryTable(table)
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

		var l []interface{}
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
						m[strings.ToLower(fname)] = val.FieldByName(fname).Interface() //输出为小写
					}
					ml = append(ml, m)
				}
			}

			//存到redis
			value, _ := json.Marshal(ml) //转成json格式存起来
			_, err := rs.Do("SET", key, value, "EX", "300")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("from db")
			return ml, nil
		}
	}
	return nil, err
}
