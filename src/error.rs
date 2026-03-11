use thiserror::Error;

pub type AppResult<T> = Result<T, AppError>;

#[derive(Debug, Error)]
pub enum AppError {
    #[error("invalid hour: {0}")]
    InvalidHour(String),
    #[error("invalid timezone: {0}")]
    InvalidTimeZone(String),
    #[error("invalid datetime: {0}")]
    InvalidDateTime(String),
    #[error("ambiguous datetime: {0}")]
    AmbiguousDateTime(String),
    #[error("nonexistent datetime: {0}")]
    NonexistentDateTime(String),
    #[error("{0}")]
    Io(#[from] std::io::Error),
}

impl AppError {
    pub fn from_local_result<T>(raw_input: &str, result: chrono::LocalResult<T>) -> Option<Self> {
        match result {
            chrono::LocalResult::None => Some(Self::NonexistentDateTime(raw_input.to_owned())),
            chrono::LocalResult::Ambiguous(_, _) => {
                Some(Self::AmbiguousDateTime(raw_input.to_owned()))
            }
            chrono::LocalResult::Single(_) => None,
        }
    }
}
