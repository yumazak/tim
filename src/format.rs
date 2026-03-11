use chrono::{DateTime, FixedOffset};

pub fn format_hour(hour: u32) -> String {
    hour.to_string()
}

pub fn format_datetime(dt: DateTime<FixedOffset>) -> String {
    dt.to_rfc3339()
}
