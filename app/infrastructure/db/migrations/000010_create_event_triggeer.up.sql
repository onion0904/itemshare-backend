-- events テーブルのトリガー
CREATE TRIGGER set_updated_at_events
BEFORE UPDATE ON events
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();