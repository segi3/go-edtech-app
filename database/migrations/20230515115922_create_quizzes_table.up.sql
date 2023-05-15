CREATE TABLE quizzes (
    `id` INT NOT NULL AUTO_INCREMENT,
    `lesson_id`INT NULL,
    `question_id`INT NULL,

    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    
    PRIMARY KEY ( `id` ),
    INDEX idx_quizzes_lesson_id ( `lesson_id` ),
    INDEX idx_quizzes_question_id ( `question_id` ),
    INDEX idx_quizzes_created_by ( `created_by` ),
    INDEX idx_quizzes_updated_by ( `updated_by` ),
    CONSTRAINT FK_quizzes_lesson_id FOREIGN KEY (`lesson_id`) REFERENCES lessons(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_quizzes_question_id FOREIGN KEY (`question_id`) REFERENCES questions(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_quizzes_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_quizzes_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`)  ON DELETE SET NULL
)