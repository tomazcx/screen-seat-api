CREATE TABLE IF NOT EXISTS movie (
	id UUID PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	duration INTEGER NOT NULL,
	poster VARCHAR(255),
	age_rating VARCHAR(5) NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS category (
	id UUID PRIMARY KEY,
	name VARCHAR(255) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS join_movie_category (
	movie_id UUID,
	category_name VARCHAR(255),
	CONSTRAINT pk_join_movie_category PRIMARY KEY(movie_id, category_name),
	CONSTRAINT fk_movie_join_movie_category FOREIGN KEY(movie_id) REFERENCES movie(id) ON DELETE CASCADE,
	CONSTRAINT fk_category_join_movie_category FOREIGN KEY(category_name) REFERENCES category(name) ON DELETE CASCADE
);
