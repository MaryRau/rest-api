CREATE TABLE Todos (
    id SERIAL PRIMARY KEY,
    title varchar(30) NOT NULL,
    description varchar(50) NOT NULL,
    isCompleted BOOLEAN NOT NULL DEFAULT FALSE,
    createdAt TIMESTAMP NOT NULL,
    completedAt TIMESTAMP
);