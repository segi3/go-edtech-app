CREATE TABLE courses (
    `id` INT NOT NULL AUTO_INCREMENT,
    `lesson_id`INT NULL,
    `product_id`INT NULL,
    `index` INT NOT NULL,

    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    
    PRIMARY KEY ( `id` ),
    INDEX idx_courses_lesson_id ( `lesson_id` ),
    INDEX idx_courses_product_id ( `product_id` ),
    INDEX idx_courses_created_by ( `created_by` ),
    INDEX idx_courses_updated_by ( `updated_by` ),
    CONSTRAINT FK_courses_lesson_id FOREIGN KEY (`lesson_id`) REFERENCES lessons(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_courses_product_id FOREIGN KEY (`product_id`) REFERENCES products(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_courses_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`)  ON DELETE SET NULL,
    CONSTRAINT FK_courses_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`)  ON DELETE SET NULL
)