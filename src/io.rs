use std::io::{self, BufRead, Write};

use crate::error::AppResult;

pub fn process_stdin_lines(
    mut process: impl FnMut(&str) -> Result<String, String>,
) -> AppResult<u8> {
    let stdin = io::stdin();
    let mut had_error = false;

    for line in stdin.lock().lines() {
        match line {
            Ok(line) => match process(&line) {
                Ok(output) => write_stdout_line(&output)?,
                Err(err) => {
                    had_error = true;
                    write_stderr_line(&err)?;
                }
            },
            Err(err) => {
                had_error = true;
                write_stderr_line(&format!("io error: {err}"))?;
                break;
            }
        }
    }

    Ok(if had_error { 1 } else { 0 })
}

pub fn write_stdout_line(line: &str) -> AppResult<()> {
    let mut stdout = io::stdout().lock();
    writeln!(stdout, "{line}")?;
    Ok(())
}

pub fn write_stderr_line(line: &str) -> AppResult<()> {
    let mut stderr = io::stderr().lock();
    writeln!(stderr, "{line}")?;
    Ok(())
}
