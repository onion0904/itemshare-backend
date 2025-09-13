-- groups テーブルのトリガー
CREATE TRIGGER set_updated_at_groups
BEFORE UPDATE ON "groups"
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();