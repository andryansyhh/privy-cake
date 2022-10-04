package query

const (
	PostCake = `
	INSERT INTO
		cake(
			title,
			description,
			rating,
			image,
			created_at,
			updated_at		
		)
		VALUES(
			?, ?, ?, ?, now(), now()
		)
	`

	GetAllCakes = `
	SELECT
		id,
		title,
		description,
		rating,
		image,
		created_at,
		updated_at	
	FROM cake WHERE deleted_at is null
	`

	GetCakeByID = `
	SELECT
		id,
		title,
		description,
		rating,
		image,
		created_at,
		updated_at	
	FROM cake 
	WHERE deleted_at is null 
	AND id = ?
	`

	UpdateCakeByID = `
	UPDATE cake SET
			title = ?,
			description = ?,
			rating = ?,
			image = ?,
			updated_at	= now()	
	WHERE id = ?
	`

	DeleteCakeByID = `
	UPDATE cake SET deleted_at = now()
	WHERE id = ?
	`
)
