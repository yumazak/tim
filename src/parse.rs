use chrono::{DateTime, FixedOffset, NaiveDateTime};

use crate::error::{AppError, AppResult};

#[derive(Debug, Clone, Copy)]
pub struct HourInput(pub u32);

impl HourInput {
    pub fn parse(input: &str) -> AppResult<Self> {
        let value = input
            .parse::<u32>()
            .map_err(|_| AppError::InvalidHour(input.to_owned()))?;
        if value <= 23 {
            Ok(Self(value))
        } else {
            Err(AppError::InvalidHour(input.to_owned()))
        }
    }
}

#[derive(Debug, Clone)]
pub enum DateTimeInput {
    Naive(NaiveDateTime),
    WithOffset(DateTime<FixedOffset>),
}

impl DateTimeInput {
    pub fn parse(input: &str) -> AppResult<Self> {
        if let Ok(value) = DateTime::parse_from_rfc3339(input) {
            return Ok(Self::WithOffset(value));
        }

        for format in [
            "%Y-%m-%dT%H:%M:%S",
            "%Y-%m-%d %H:%M:%S",
            "%Y-%m-%dT%H:%M",
            "%Y-%m-%d %H:%M",
        ] {
            if let Ok(value) = NaiveDateTime::parse_from_str(input, format) {
                return Ok(Self::Naive(value));
            }
        }

        Err(AppError::InvalidDateTime(input.to_owned()))
    }
}
