CREATE TABLE lessons (
    `id` INT NOT NULL AUTO_INCREMENT,
    
    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `video_content` VARCHAR(255) NOT NULL,
    `text_content` TEXT NOT NULL,

    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY ( `id` ),
    INDEX idx_lessons_created_by ( `created_by` ) ,
    INDEX idx_lessons_updated_by ( `updated_by` ) ,
    CONSTRAINT FK_lesson_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_lesson_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`)  ON DELETE SET NULL
)