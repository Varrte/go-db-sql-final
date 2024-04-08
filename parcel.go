package main

import (
	"database/sql"
	"fmt"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {
	// реализуйте добавление строки в таблицу parcel, используйте данные из переменной p

	res, err := s.db.Exec("INSERT INTO parsel(client,status,address,created_at) VALUES (:client,:status,:address,:created_at)",
		sql.Named("client", p.Client),
		sql.Named("status", p.Status),
		sql.Named("address", p.Address),
		sql.Named("created_at", p.CreatedAt))
	id, err := res.LastInsertId()
	// верните идентификатор последней добавленной записи
	return int(id), err
}

func (s ParcelStore) Get(number int) (Parcel, error) {
	// реализуйте чтение строки по заданному number
	// здесь из таблицы должна вернуться только одна строка
	row := s.db.QueryRow("SELECT client,status,address,created_at FROM parsel WHERE id=:number",
		sql.Named("number", number))
	//defer row.Close()

	// заполните объект Parcel данными из таблицы
	p := Parcel{}
	err := row.Scan(&p.Client, &p.Status, &p.Address, &p.CreatedAt)

	return p, err
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	// реализуйте чтение строк из таблицы parcel по заданному client
	// здесь из таблицы может вернуться несколько строк
	rows, err := s.db.Query("SELECT client,status,address,created_at FROM parsel WHERE id=:number",
		sql.Named("number", client))

	// заполните срез Parcel данными из таблицы
	var res []Parcel
	i := 0
	for rows.Next() {

		err := rows.Scan(res[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		i++
	}
	return res, err
}

func (s ParcelStore) SetStatus(number int, status string) error {
	// реализуйте обновление статуса в таблице parcel
	_, err := s.db.Exec("UPDATE parsel SET status =:status WHERE id=:id)",
		sql.Named("status", status),
		sql.Named("id", number))

	return err
}

func (s ParcelStore) SetAddress(number int, address string) error {
	// реализуйте обновление адреса в таблице parcel
	// менять адрес можно только если значение статуса registered

	return nil
}

func (s ParcelStore) Delete(number int) error {
	// реализуйте удаление строки из таблицы parcel
	// удалять строку можно только если значение статуса registered

	return nil
}
