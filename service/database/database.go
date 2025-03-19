/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	CreateUser(u User) (User, error)

	GetUserByName(username string) (User, error)

	GetGroupByID(groupID int) (Group, error)

	GetUserByID(userID int) (User, error)

	ExistName(username string) (bool, error)

	ChangeUsername(userID int, newUsername string) error

	CreateGroup(g Group, userID int) (Group, error)

	ExistGroup(groupID int) (bool, error)

	ExistUserInGroup(userID int, groupID int) (bool, error)

	AddUserToGroup(userID int, groupID int) error

	DeleteUserFromGroup(groupID int, userID int) error

	DeleteGroup(groupID int) error

	GetMemberGroup(groupID int) ([]int, error)

	ChangeGroupName(groupID int, newGroupName string) error

	CreateConversation(c Conversation) (Conversation, error)

	CreateMessage(m Message) (Message, error)

	DeleteUser(userID int) error

	SearchUser(usersearch string) ([]User, error)

	GetConversation(conversationID int) (Conversation, error)

	GetListConversations(userID int) ([]Conversation, error)

	AddMemberPrivate(conversationID int, userID int) error

	ExistUserInConv(userID int, conversationID int) (bool, error)

	GetUserInConversationPrivate(conversationID int, userID int) (User, error)

	GetMessage(conversationID int, messageID int) (Message, error)

	ExistUserID(userID int) (bool, error)

	ExistConversation(userID1 int, userID2 int) (bool, error)

	ExistMessage(messageID int, conversationID int) (bool, error)

	CreateComment(comment Comment) (Comment, error)

	DeleteMessage(messageID int, conversationID int) error

	DeleteComment(commentID int) error

	ExistComment(commentID int, messageID int) (bool, error)

	GetAllMessage(conversationID int) ([]Message, error)

	GetUsersInGroup(groupID int) ([]int, error)

	ExistConversationByID(conversationID int) (bool, error)

	GetComments(messageID int, conversationID int) ([]Comment, error)

	GetExistComment(messageID int, conversationID int, userID int) (int, error)

	UpdateComment(commentID int, commentTXT string) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	var tableCount uint8
	err := db.QueryRow(`SELECT count(name) FROM sqlite_master WHERE type='table';`).Scan(&tableCount)
	if err != nil {
		return nil, fmt.Errorf("error checking if database is empty: %w", err)
	}

	if tableCount != 7 {
		// -----CREATE USER TABLE ------- //
		_, err = db.Exec(sql_USERTABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)

		}

		_, err = db.Exec(sql_GROUPTABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)

		}

		_, err = db.Exec(sql_MEMBERGROUPTABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)

		}

		_, err = db.Exec(sqL_CONVERSATIONTABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)

		}

		_, err = db.Exec(sql_MESSAGETABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)

		}

		_, err = db.Exec(sql_MEMBERPRIVATETABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)

		}

		_, err = db.Exec(sql_COMMENTTABLE)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)

		}

	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	/*var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}*/

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
