use crate::cli::{Cli, Command};
use crate::commands;
use crate::error::AppResult;

pub fn run(cli: Cli) -> AppResult<u8> {
    match cli.command {
        Command::H(args) => commands::h::run(args),
        Command::Dt(args) => commands::dt::run(args),
        Command::Tz => commands::tz_list::run(),
    }
}
