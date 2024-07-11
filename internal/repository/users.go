package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"time-tracker/internal/model"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) GetUsers(filters map[string]string, limit int, cursor int) ([]model.User, error) {
	var users []model.User

	query := `SELECT id, passport_number, name, surname, patronymic, address FROM users WHERE id > $1`
	args := []interface{}{cursor}

	for field, value := range filters {
		query += fmt.Sprintf(" AND %s = $%d", field, len(args)+1)
		args = append(args, value)
	}

	if limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", len(args)+1)
		args = append(args, limit)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

func (r *UsersRepo) CreateUser(user model.User) (userId int, err error) {

	query := `INSERT INTO users (passport_number, name, surname, patronymic, address) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	er := r.db.QueryRow(query, user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address).Scan(&userId)
	if er != nil {
		return 0, fmt.Errorf("failed to create user: %w", er)
	}
	return userId, nil
}

func (r *UsersRepo) DeleteUser(userId int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, userId)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d does not exist", userId)
	}
	return nil
}

func (r *UsersRepo) UpdateUser(userId int, user model.UpdateUserInput) error {

	query := `UPDATE users SET `

	rt := reflect.TypeOf(user)
	rv := reflect.ValueOf(user)
	var values []interface{}
	queryLastIndex := 1

	for i := 0; i < rv.NumField(); i++ {
		if !rv.Field(i).IsZero() {

			values = append(values, rv.Field(i).Interface())
			query += rt.Field(i).Tag.Get("sql") + " = $" + strconv.Itoa(queryLastIndex) + ", "
			queryLastIndex++
		}
	}

	if len(values) == 0 {
		return fmt.Errorf("empty request body")
	}

	query = query[:len(query)-2]
	query += " WHERE id = $" + strconv.Itoa(queryLastIndex)
	values = append(values, userId)
	result, err := r.db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d does not exist", userId)
	}
	return nil
}

func (r *UsersRepo) GetUserById(userId int) (model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userId).Scan(&user.Id, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, fmt.Errorf("user with ID %d does not exist", userId)
		}
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}
