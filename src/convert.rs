use chrono::{DateTime, FixedOffset, NaiveDateTime, Offset, TimeZone};
use chrono_tz::Tz;

use crate::error::{AppError, AppResult};
use crate::parse::{DateTimeInput, HourInput};
use crate::tz::fixed_offset_for_hour;

pub fn convert_hour(input: HourInput, from: Tz, to: Tz) -> u32 {
    let from_secs = fixed_offset_for_hour(from).local_minus_utc();
    let to_secs = fixed_offset_for_hour(to).local_minus_utc();
    let total_minutes = (input.0 as i32) * 60 + (to_secs - from_secs) / 60;
    total_minutes.div_euclid(60).rem_euclid(24) as u32
}

pub fn convert_datetime(
    input: DateTimeInput,
    from: Tz,
    to: Tz,
    raw_input: &str,
) -> AppResult<DateTime<FixedOffset>> {
    let source = match input {
        DateTimeInput::WithOffset(value) => value,
        DateTimeInput::Naive(value) => resolve_naive_datetime(value, from, raw_input)?,
    };

    let to_offset = to.offset_from_utc_datetime(&source.naive_utc()).fix();
    Ok(source.with_timezone(&to_offset))
}

fn resolve_naive_datetime(
    value: NaiveDateTime,
    from: Tz,
    raw_input: &str,
) -> AppResult<DateTime<FixedOffset>> {
    match from.from_local_datetime(&value) {
        chrono::LocalResult::Single(result) => Ok(result.with_timezone(&result.offset().fix())),
        other => {
            Err(AppError::from_local_result(raw_input, other).expect("non-single local result"))
        }
    }
}
