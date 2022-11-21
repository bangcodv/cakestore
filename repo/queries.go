package repo

const (
	QueryGetListCake = `select id,title,description,rating,image,created_at, updated_at 
		from cake 
		where is_active=1 limit ? offset ?;`

	QueryGetDetailCake = `select id,title,description,rating,image,created_at, updated_at 
		from cake 
		where id=? and is_active=1 
		limit 1;
	`

	QueryInsertCake = `INSERT INTO cake(title,description,rating,image) VALUES(?,?,?,?)`

	QueryUpdateCake = `update cake
	set title=?, description=?, rating = ?, image=?, updated_at=now()
	where id=?;
	`
	QuerySoftDeleteCake = `update cake
	set is_active=0, updated_at=now()
	where id=?;
	`
)
