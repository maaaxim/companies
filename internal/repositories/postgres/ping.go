package postgres

func (r *Repository) Ping() bool {
	err := r.DB.Ping()

	return err == nil
}
