use crate::error::AppResult;
use crate::io::write_stdout_line;
use crate::tz::all_timezones;

pub fn run() -> AppResult<u8> {
    for tz in all_timezones() {
        write_stdout_line(&tz.to_string())?;
    }
    Ok(0)
}
