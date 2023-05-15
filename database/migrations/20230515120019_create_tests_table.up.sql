CREATE TABLE tests (
    `id` INT NOT NULL AUTO_INCREMENT,
    `question_id`INT NULL,
    `user_id`INT NULL,

    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    
    PRIMARY KEY ( `id` ),
    INDEX idx_tests_question_id ( `question_id` ),
    INDEX idx_tests_user_id ( `user_id` ),
    INDEX idx_tests_created_by ( `created_by` ),
    INDEX idx_tests_updated_by ( `updated_by` ),
    CONSTRAINT FK_tests_question_id FOREIGN KEY (`question_id`) REFERENCES questions(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_tests_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_tests_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_tests_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`)  ON DELETE SET NULL
)