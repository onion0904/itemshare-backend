-- ユーザーが存在しなければ作成
CREATE USER IF NOT EXISTS 'app_user'@'%' IDENTIFIED BY 'password';

-- 念のため全権限をはく奪（エラー防止のためユーザー存在確認後）
REVOKE ALL PRIVILEGES ON *.* FROM 'app_user'@'%';

-- 必要な権限だけ付与
GRANT SELECT, INSERT, UPDATE, DELETE ON myapp_db.* TO 'app_user'@'%';

-- 権限の反映
FLUSH PRIVILEGES;
