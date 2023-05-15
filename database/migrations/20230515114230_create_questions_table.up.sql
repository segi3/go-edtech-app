CREATE TABLE questions (
    `id` INT NOT NULL AUTO_INCREMENT,

    `thumbnail` VARCHAR(255) NULL,
    
    `question` VARCHAR(255) NOT NULL,
    `option_a` VARCHAR(255) NOT NULL,
    `option_b` VARCHAR(255) NOT NULL,
    `option_c` VARCHAR(255) NOT NULL,
    `correct_option` INT NOT NULL,

    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,

    PRIMARY KEY ( `id` ),
    INDEX idx_questions_created_by ( `created_by` ) ,
    INDEX idx_questions_updated_by ( `updated_by` ) ,
    CONSTRAINT FK_question_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_question_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`)  ON DELETE SET NULL
) 