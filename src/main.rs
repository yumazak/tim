mod app;
mod cli;
mod commands;
mod convert;
mod error;
mod format;
mod io;
mod parse;
mod tz;

use std::io::ErrorKind;
use std::process::ExitCode;

fn main() -> ExitCode {
    let cli = cli::Cli::parse_args();

    match app::run(cli) {
        Ok(code) => ExitCode::from(code),
        Err(err) => {
            if is_broken_pipe(&err) {
                return ExitCode::from(0);
            }
            eprintln!("{err}");
            ExitCode::from(1)
        }
    }
}

fn is_broken_pipe(err: &error::AppError) -> bool {
    matches!(err, error::AppError::Io(e) if e.kind() == ErrorKind::BrokenPipe)
}
