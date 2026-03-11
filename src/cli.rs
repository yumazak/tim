use clap::{Args, Parser, Subcommand};

#[derive(Debug, Parser)]
#[command(name = "tim")]
#[command(about = "Time zone conversion CLI")]
pub struct Cli {
    #[command(subcommand)]
    pub command: Command,
}

#[derive(Debug, Subcommand)]
pub enum Command {
    /// Convert hour only
    H(HCommand),
    /// Convert datetime
    Dt(DateTimeCommand),
    /// List all IANA time zones
    Tz,
}

#[derive(Debug, Args, Clone)]
pub struct ZoneArgs {
    #[arg(short = 'f', long = "from", default_value = "Asia/Tokyo")]
    pub from: String,
    #[arg(short = 't', long = "to", default_value = "UTC")]
    pub to: String,
}

#[derive(Debug, Args)]
pub struct HCommand {
    #[command(flatten)]
    pub zones: ZoneArgs,
    pub hour: Option<String>,
}

#[derive(Debug, Args)]
pub struct DateTimeCommand {
    #[command(flatten)]
    pub zones: ZoneArgs,
    pub datetime: Option<String>,
}

impl Cli {
    pub fn parse_args() -> Self {
        Self::parse()
    }
}
