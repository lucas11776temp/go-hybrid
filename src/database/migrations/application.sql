CREATE TABLE `users` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE `courses` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `title` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL
);

CREATE TABLE `lessons` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `course_id` integer NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `order` integer NULL,
  `title` VARCHAR(255) NOT NULL,
  `notes` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL
);

INSERT INTO courses(title, description) VALUES("Learn Go Days", "Short course to learn GO in a go.");
INSERT INTO courses(title, description) VALUES("System Design", "We will learn system design in days.");