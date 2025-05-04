package time

import "time"

func Now() time.Time {
	return time.Now()
}

//次の週の日曜日の日を出力
func NextEndWeek() time.Time {
	now := time.Now()
	// 現在の日付から次の日曜日までの日数を計算
	daysUntilSunday := (7 - int(now.Weekday())) % 7
	// 次の週の日曜日までさらに7日を加算
	return now.AddDate(0, 0, daysUntilSunday+7)
}

func NextStartWeek() time.Time {
	now := time.Now()
    daysUntilSunday := (7 - int(now.Weekday())) % 7
    return now.AddDate(0, 0, daysUntilSunday+1)
}

func CreateEventDate(year, month, day int32) time.Time {
    return time.Date(
        int(year),              // year は int に変換
        time.Month(month),      // month は time.Month に変換
        int(day),               // day も int に変換
        0, 0, 0, 0,             // 時間（00:00:00.000）
        time.UTC,               // タイムゾーン
    )
}