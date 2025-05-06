-- 1) ロール（ユーザー）を存在しなければ作成
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT FROM pg_catalog.pg_roles
    WHERE rolname = 'app_user'
  ) THEN
    CREATE ROLE app_user
      LOGIN
      PASSWORD 'password';
  END IF;
END
$$;

-- 2) いったん既存の権限を剥奪（エラー防止のため）
REVOKE ALL PRIVILEGES
  ON DATABASE myapp_db
  FROM app_user;
REVOKE ALL PRIVILEGES
  ON SCHEMA public
  FROM app_user;
REVOKE ALL PRIVILEGES
  ON ALL TABLES IN SCHEMA public
  FROM app_user;

-- 3) 必要な権限だけを付与
GRANT CONNECT
  ON DATABASE myapp_db
  TO app_user;
GRANT USAGE
  ON SCHEMA public
  TO app_user;
GRANT SELECT, INSERT, UPDATE, DELETE
  ON ALL TABLES IN SCHEMA public
  TO app_user;

-- （オプション）将来作成されるテーブルにも同じ権限をデフォルトで付与
ALTER DEFAULT PRIVILEGES
  IN SCHEMA public
  GRANT SELECT, INSERT, UPDATE, DELETE
  ON TABLES
  TO app_user;
