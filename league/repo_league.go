package league

// RepoLeague struct
type RepoLeague struct {
	ID          int64  `db:"id"`
	UniqueName  string `db:"unique_name"`
	DisplayName string `db:"display_name"`
}
