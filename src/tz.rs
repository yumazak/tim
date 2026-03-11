use chrono::{FixedOffset, NaiveDate, Offset, TimeZone};
use chrono_tz::{TZ_VARIANTS, Tz};

use crate::error::{AppError, AppResult};

pub fn parse_tz(input: &str) -> AppResult<Tz> {
    input
        .parse::<Tz>()
        .map_err(|_| AppError::InvalidTimeZone(input.to_owned()))
}

pub fn fixed_offset_for_hour(tz: Tz) -> FixedOffset {
    let epoch = NaiveDate::from_ymd_opt(1970, 1, 1)
        .unwrap()
        .and_hms_opt(0, 0, 0)
        .unwrap();
    tz.offset_from_utc_datetime(&epoch).fix()
}

pub fn all_timezones() -> &'static [Tz] {
    &TZ_VARIANTS
}
