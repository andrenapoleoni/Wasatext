package database

// -------USER Table------- //

var sql_USERTABLE = ` CREATE TABLE IF NOT EXISTS User
(
	userID INTEGER NOT NULL UNIQUE,
	username STRING NOT NULL UNIQUE,
	PRIMARY KEY(userID)	
);`

// ---------MESSAGE TABLE-------- //
var sql_MESSAGETABLE = ` CREATE TABLE IF NOT EXISTS Message
(
	conversationID INTEGER NOT NULL ,
	userID INTEGER NOT NULL ,
	messageID INTEGER NOT NULL,
	message TEXT NOT NULL,
	TIMESTAMP DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(messageID)
);`

// -------CONVERSATION TABLE------- //

var sqL_CONVERSATIONTABLE = ` CREATE TABLE IF NOT EXISTS Conversation 
(
	conversationID INTEGER NOT NULL,
	groupID INTEGER,
	PRIMARY KEY(conversationID)

	


);`

// ------GROUP TABLE------- //
var sql_GROUPTABLE = ` CREATE TABLE IF NOT EXISTS Groupchat 
(
	groupID INTEGER NOT NULL UNIQUE,
	groupName TEXT NOT NULL,
	PRIMARY KEY(groupID)


);`

// --------MEMBER GROUP TABLE------- //
var sql_MEMBERGROUPTABLE = ` CREATE TABLE IF NOT EXISTS MemberGroup 
(
	groupID INTEGER,
	userID INTEGER NOT NULL,
	CONSTRAINT usersMember
		FOREIGN KEY(userID) REFERENCES User(userID)
		ON DELETE CASCADE


);`

// -----------MEMBER PRIVATE CHAT-------- //
var sql_MEMBERPRIVATETABLE = ` CREATE TABLE IF NOT EXISTS MemberPrivate
(
	conversationID INTEGER NOT NULL,
	userID INTEGER NOT NULL,
	CONSTRAINT memberPrivate
		FOREIGN KEY(conversationID) REFERENCES Conversation(conversationID)

);`

// -------COMMENT TABLE------- //
var sql_COMMENTTABLE = ` CREATE TABLE IF NOT EXISTS Comment
(
	commentID INTEGER NOT NULL,
	messageID INTEGER NOT NULL,
	commentTXT TEXT NOT NULL,
	conversationID INTEGER NOT NULL,
	userID INTEGER NOT NULL,
	PRIMARY KEY(commentID),
	FOREIGN KEY(messageID) REFERENCES Message(messageID) ON DELETE CASCADE
	
);`
