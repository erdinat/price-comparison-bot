package repository

import "database/sql"

type Site struct {
	ID   int
	Name string
	URL  string
}

type SiteRepository struct {
	DB *sql.DB
}

func NewSiteRepository(db *sql.DB) *SiteRepository {
	return &SiteRepository{DB: db}
}

func (r *SiteRepository) GetAllSites() ([]Site, error) {
	rows, err := r.DB.Query("SELECT * FROM sites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []Site
	for rows.Next() {
		var site Site
		if err := rows.Scan(&site.ID, &site.Name, &site.URL); err != nil {
			return nil, err
		}

		sites = append(sites, site)
	}

	return sites, nil

}
