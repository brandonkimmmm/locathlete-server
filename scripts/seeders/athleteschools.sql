INSERT INTO "athlete_schools" (
	athlete_id,
	school_id,
	start_date,
	end_date,
	created_at
) VALUES (
	(SELECT id FROM "athletes" WHERE last_name = 'Westbrook'),
	(SELECT id FROM "schools" WHERE name = 'Leuzinger High School'),
	'09-01-2004',
	'06-01-2006',
	NOW()
),
(
	(SELECT id FROM "athletes" WHERE last_name = 'Harden'),
	(SELECT id FROM "schools" WHERE name = 'Artesia High School'),
	'09-01-2004',
	'06-01-2007',
	NOW()
),
(
	(SELECT id FROM "athletes" WHERE last_name = 'DeRozan'),
	(SELECT id FROM "schools" WHERE name = 'Compton High School'),
	'09-01-2005',
	'06-01-2008',
	NOW()
),
(
	(SELECT id FROM "athletes" WHERE last_name = 'Leonard'),
	(SELECT id FROM "schools" WHERE name = 'Martin Luther King High School'),
	'09-01-2007',
	'06-01-2009',
	NOW()
),
(
	(SELECT id FROM "athletes" WHERE last_name = 'Ball'),
	(SELECT id FROM "schools" WHERE name = 'Chino Hills High School'),
	'09-01-2015',
	'06-01-2017',
	NOW()
),
(
	(SELECT id FROM "athletes" WHERE last_name = 'George'),
	(SELECT id FROM "schools" WHERE name = 'Knight High School'),
	'09-01-2005',
	'06-01-2008',
	NOW()
)