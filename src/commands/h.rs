use chrono_tz::Tz;

use crate::cli::HCommand;
use crate::convert::convert_hour;
use crate::error::AppResult;
use crate::format::format_hour;
use crate::io::{process_stdin_lines, write_stdout_line};
use crate::parse::HourInput;
use crate::tz::parse_tz;

pub fn run(args: HCommand) -> AppResult<u8> {
    let from = parse_tz(&args.zones.from)?;
    let to = parse_tz(&args.zones.to)?;

    match args.hour {
        Some(hour) => {
            let output = process_one(&hour, from, to)?;
            write_stdout_line(&output)?;
            Ok(0)
        }
        None => process_stdin_lines(|line| process_one(line, from, to).map_err(|e| e.to_string())),
    }
}

fn process_one(input: &str, from: Tz, to: Tz) -> AppResult<String> {
    let hour = HourInput::parse(input.trim())?;
    Ok(format_hour(convert_hour(hour, from, to)))
}
