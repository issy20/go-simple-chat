-- CREATE TABLE IF NOT EXISTS `members` (
--   `room_id` int(11) NOT NULL,
--   `user_id` int(11) NOT NULL,
--   `number` int(11),
--   PRIMARY KEY(room_id, user_id),
--   FOREIGN KEY (room_id) REFERENCES `rooms`(id),
--   FOREIGN KEY (user_id) REFERENCES `users`(id)
-- ) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- ALTER TABLE `members` DROP FOREIGN KEY members_ibfk_1;
-- ALTER TABLE `members` DROP FOREIGN KEY members_ibfk_2;

-- ALTER TABLE `members` ADD CONSTRAINT `members_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;
-- ALTER TABLE `members` ADD CONSTRAINT `members_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;

CREATE TABLE IF NOT EXISTS `members` (
  `room_id` INT NOT NULL,
  `user_id` INT NOT NULL,
  PRIMARY KEY(room_id, user_id),
  FOREIGN KEY (room_id) REFERENCES `rooms`(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES `users`(id) ON DELETE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;