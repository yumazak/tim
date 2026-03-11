use chrono_tz::Tz;

use crate::cli::DateTimeCommand;
use crate::convert::convert_datetime;
use crate::error::AppResult;
use crate::format::format_datetime;
use crate::io::{process_stdin_lines, write_stdout_line};
use crate::parse::DateTimeInput;
use crate::tz::parse_tz;

pub fn run(args: DateTimeCommand) -> AppResult<u8> {
    let from = parse_tz(&args.zones.from)?;
    let to = parse_tz(&args.zones.to)?;

    match args.datetime {
        Some(datetime) => {
            let output = process_one(&datetime, from, to)?;
            write_stdout_line(&output)?;
            Ok(0)
        }
        None => process_stdin_lines(|line| process_one(line, from, to).map_err(|e| e.to_string())),
    }
}

fn process_one(input: &str, from: Tz, to: Tz) -> AppResult<String> {
    let trimmed = input.trim();
    let datetime = DateTimeInput::parse(trimmed)?;
    let converted = convert_datetime(datetime, from, to, trimmed)?;
    Ok(format_datetime(converted))
}
