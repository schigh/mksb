SELECT
  u.`id`                                        AS 'id',
  u.`name`                                      AS 'name',
  u.`email`                                     AS 'email',
  IF(IFNULL(a.`accountID`, 0) > 0, b'1', b'0')  AS 'hasAccount',
  IF(u.`isAnonymous` > 0, b'1', b'0')           AS 'isAnonymous',
  IF(IFNULL(us.`userID`, 0) > 0, b'1', b'0')    AS 'isV7Signup',
  u.`avatarID`                                  AS 'avatarID',
  IFNULL(a.`vendorID`, '')                      AS 'vendorID',
  IFNULL(av.`assetKey`, '')                     AS 'assetKey',
  u.`createdAt`                                 AS 'createdAt',
  u.`updatedAt`                                 AS 'updatedAt'
FROM `user` u
       LEFT JOIN `account` a ON u.`id` = a.`userID`
       LEFT JOIN `user_v7_avatars` av ON u.`id` = av.`userID`
       LEFT JOIN `user_v7_signups` us ON u.`id` = us.`userID`
WHERE
  u.`id` = ?
