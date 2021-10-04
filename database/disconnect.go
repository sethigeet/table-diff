package database

func Disconnect() error {
	err := DB.Close()

	return err
}
