package db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type TbCommitHash struct {
	Id         int64     `orm:"column(id);pk"`
	BlockId    int64     `orm:"column(block_id);null" description:"区块号"`
	BlockHash  string    `orm:"column(block_hash);size(255);null" description:"blockHash"`
	TxTime     time.Time `orm:"column(tx_time);type(datetime);null" description:"交易时间"`
	TxHash     string    `orm:"column(tx_hash);size(255);null" description:"交易hash"`
	Commiter   string    `orm:"column(commiter);size(255);null" description:"提交者地址"`
	CommitHash string    `orm:"column(commithash);size(255);null" description:"Commit hash"`
	Status     int       `orm:"column(status);null" description:"状态,1-正常 2-冻结"`
	IsDeleted  int8      `orm:"column(is_deleted);null" description:"删除状态 0-正常 1-删除"`
	SyncTime   time.Time `orm:"column(sync_time);type(timestamp);null;auto_now_add" description:"同步时间"`
	CreateTime time.Time `orm:"column(create_time);type(timestamp);null;auto_now_add" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp);null;auto_now_add" description:"更新时间"`
}

func (t *TbCommitHash) TableName() string {
	return "tb_commithash"
}

func init() {
	orm.RegisterModel(new(TbCommitHash))
}

// AddTbCommitHash insert a new TbCommitHash into database and returns
// last inserted Id on success.
func AddTbCommitHash(m *TbCommitHash) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTbCommitHashById retrieves TbCommitHash by Id. Returns error if
// Id doesn't exist
func GetTbCommitHashById(id int64) (v *TbCommitHash, err error) {
	o := orm.NewOrm()
	v = &TbCommitHash{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTbCommitHash retrieves all TbCommitHash matches certain condition. Returns empty list if
// no records exist
func GetAllTbCommitHash(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TbCommitHash))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
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

	var l []TbCommitHash
	qs = qs.OrderBy(sortFields...)
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
		return ml, nil
	}
	return nil, err
}

// UpdateTbCommitHash updates TbCommitHash by Id and returns error if
// the record to be updated doesn't exist
func UpdateTbCommitHashById(m *TbCommitHash) (err error) {
	o := orm.NewOrm()
	v := TbCommitHash{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTbCommitHash deletes TbCommitHash by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTbCommitHash(id int64) (err error) {
	o := orm.NewOrm()
	v := TbCommitHash{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TbCommitHash{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
