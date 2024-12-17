CREATE TABLE IF NOT EXISTS patterns (id INTEGER PRIMARY KEY AUTOINCREMENT,
                                        name TEXT NOT NULL UNIQUE,
                                        description TEXT,
                                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE IF NOT EXISTS problems (id INTEGER PRIMARY KEY AUTOINCREMENT,
                                        pattern_id INTEGER NOT NULL,
                                        name TEXT NOT NULL,
                                        description TEXT,
                                        solution TEXT,
                                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                                        FOREIGN KEY (pattern_id) REFERENCES patterns(id),
                                        UNIQUE (pattern_id, name));
