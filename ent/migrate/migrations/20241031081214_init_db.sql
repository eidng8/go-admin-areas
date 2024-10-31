-- Create "admin_areas" table
CREATE TABLE `admin_areas`
(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `deleted_at` TIMESTAMP NULL,
    `name` VARCHAR(255) NOT NULL COMMENT "Administrative area name",
    `abbr` VARCHAR(255) NULL COMMENT "Administrative area abbreviation, CSV values",
    `memo` LONGTEXT NULL COMMENT "Remarks",
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `parent_id` INT UNSIGNED NULL,
    PRIMARY KEY (`id`),
    INDEX `admin_areas_admin_areas_children` (`parent_id`),
    CONSTRAINT `admin_areas_admin_areas_children` FOREIGN KEY (`parent_id`) REFERENCES `admin_areas` (`id`) ON UPDATE NO ACTION ON DELETE RESTRICT
) CHARSET `utf8mb4`
    COLLATE `utf8mb4_unicode_ci` COMMENT "Administrative area table";
