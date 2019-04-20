package database

import (
	u "github.com/XiaoXice/AirNet/common/util"
)

type NodeList struct {
	List map[HashInfo]Node
}

func (d *DataBase) LoadNodeList() (*NodeList, error) {
	rows, err := d.SelectALlNodeList()
	if err != nil {
		return nil, err
	}
	list := make(map[HashInfo]Node)
	for _,row := range rows {
		node,err := row.ToNode()
		if err != nil {
			return nil, err
		}
		list[node.HashInfo] = *node
	}
	return &NodeList{list}, nil
}

func (d *DataBase) SaveNodeList(nl *NodeList) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("REPLACE INTO NodeList VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		d.Logger.Handler(err.Error())
		return err
	}
	for _, Node := range nl.List {
		n, err := Node.ToNodeListNew()
		if err != nil {
			return err
		}
		_, err = stmt.Exec(u.Map2Array(u.Strict2Map(*n))...)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}