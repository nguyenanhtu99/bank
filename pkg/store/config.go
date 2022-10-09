package store

type Config struct {
	MongoUrl		string	`default:"mongodb://127.0.0.1:27017" split_words:"true"`
	MongoDatabase	string	`default:"bank" split_words:"true"`
}